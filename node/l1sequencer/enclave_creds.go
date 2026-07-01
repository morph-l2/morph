package l1sequencer

import (
	"encoding/binary"
	"encoding/xml"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// awsCreds are the temporary AWS credentials the node injects into the
// enclave over vsock (op=ProvideCredentials). The enclave reads no IMDS,
// so the host resolves credentials and pushes them in.
type awsCreds struct {
	accessKeyID     string
	secretAccessKey string
	sessionToken    string
	expiryUnixSecs  uint64
}

// injectCredentials resolves credentials from the pod's IRSA environment
// and pushes them into the enclave on a one-shot connection, so the
// signer can run its deferred one-time key load (SM Get + KMS Decrypt).
//
// No IRSA env (local/dev) → skip; the following probe() then surfaces the
// enclave's own error if it actually needed credentials. Idempotent
// enclave-side: re-injecting after the key is loaded is a no-op Ack, so
// this runs safely on every node (re)start.
func (s *EnclaveSigner) injectCredentials() error {
	creds, available, err := irsaCreds()
	if err != nil {
		return fmt.Errorf("resolve IRSA credentials: %w", err)
	}
	if !available {
		s.logger.Info("no IRSA env; skipping enclave credential injection")
		return nil
	}
	conn, err := s.dial()
	if err != nil {
		return err
	}
	defer conn.Close()
	if err := conn.SetDeadline(time.Now().Add(credLoadTimeout)); err != nil {
		return fmt.Errorf("set deadline: %w", err)
	}
	if err := writeProvideCredentials(conn, creds); err != nil {
		return fmt.Errorf("provide credentials: %w", err)
	}
	s.logger.Info("enclave credentials injected via IRSA", "role", os.Getenv("AWS_ROLE_ARN"))
	return nil
}

// irsaCreds resolves temporary credentials from the EKS IRSA environment
// (AWS_ROLE_ARN + AWS_WEB_IDENTITY_TOKEN_FILE, injected by the Pod
// Identity Webhook). The token file is a projected OIDC JWT — not AWS
// credentials — so it is exchanged for a role session via STS
// AssumeRoleWithWebIdentity, the one STS call that needs no SigV4
// signing. Done with the stdlib to avoid pulling in aws-sdk-go, mirroring
// the ops-cli reference client.
//
// available=false (nil error) when the IRSA env is absent.
func irsaCreds() (creds awsCreds, available bool, err error) {
	roleARN := os.Getenv("AWS_ROLE_ARN")
	tokenFile := os.Getenv("AWS_WEB_IDENTITY_TOKEN_FILE")
	if roleARN == "" || tokenFile == "" {
		return awsCreds{}, false, nil
	}
	jwt, err := os.ReadFile(tokenFile)
	if err != nil {
		return awsCreds{}, true, fmt.Errorf("read web identity token %s: %w", tokenFile, err)
	}
	session := os.Getenv("AWS_ROLE_SESSION_NAME")
	if session == "" {
		session = "morph-enclave-signer"
	}
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = os.Getenv("AWS_DEFAULT_REGION")
	}
	creds, err = assumeRoleWithWebIdentity(region, roleARN, session, strings.TrimSpace(string(jwt)))
	return creds, true, err
}

// assumeRoleWithWebIdentity exchanges an OIDC JWT for temporary AWS
// credentials via the (unsigned) STS AssumeRoleWithWebIdentity action.
// Regional endpoint by default (matches EKS AWS_STS_REGIONAL_ENDPOINTS);
// AWS_ENDPOINT_URL_STS overrides.
func assumeRoleWithWebIdentity(region, roleARN, session, jwt string) (awsCreds, error) {
	endpoint := "https://sts.amazonaws.com/"
	if region != "" {
		endpoint = fmt.Sprintf("https://sts.%s.amazonaws.com/", region)
	}
	if v := os.Getenv("AWS_ENDPOINT_URL_STS"); v != "" {
		endpoint = v
	}
	form := url.Values{
		"Action":           {"AssumeRoleWithWebIdentity"},
		"Version":          {"2011-06-15"},
		"RoleArn":          {roleARN},
		"RoleSessionName":  {session},
		"WebIdentityToken": {jwt},
	}
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.PostForm(endpoint, form)
	if err != nil {
		return awsCreds{}, fmt.Errorf("STS request: %w", err)
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if resp.StatusCode != http.StatusOK {
		return awsCreds{}, fmt.Errorf("STS %d: %s", resp.StatusCode, strings.TrimSpace(string(raw)))
	}

	var parsed struct {
		Result struct {
			Credentials struct {
				AccessKeyID     string `xml:"AccessKeyId"`
				SecretAccessKey string `xml:"SecretAccessKey"`
				SessionToken    string `xml:"SessionToken"`
				Expiration      string `xml:"Expiration"`
			} `xml:"Credentials"`
		} `xml:"AssumeRoleWithWebIdentityResult"`
	}
	if err := xml.Unmarshal(raw, &parsed); err != nil {
		return awsCreds{}, fmt.Errorf("parse STS response: %w", err)
	}
	c := parsed.Result.Credentials
	if c.AccessKeyID == "" || c.SecretAccessKey == "" {
		return awsCreds{}, fmt.Errorf("STS response missing credentials")
	}
	out := awsCreds{accessKeyID: c.AccessKeyID, secretAccessKey: c.SecretAccessKey, sessionToken: c.SessionToken}
	if c.Expiration != "" {
		t, perr := time.Parse(time.RFC3339, c.Expiration)
		if perr != nil {
			return awsCreds{}, fmt.Errorf("STS bad Expiration %q: %w", c.Expiration, perr)
		}
		out.expiryUnixSecs = uint64(t.Unix())
	}
	return out, nil
}

// writeProvideCredentials sends op=ProvideCredentials on an established
// conn and reads the Ack. Field layout mirrors crates/protocol (and
// ops-cli's client.go):
//
//	[op=0x04][akid_lp][sak_lp][token_lp][expiry:u64]
//
// where each _lp is [len:u16][bytes]; all integers are big-endian.
func writeProvideCredentials(conn net.Conn, c awsCreds) error {
	buf := []byte{opProvideCredentials}
	buf = appendLPString(buf, c.accessKeyID)
	buf = appendLPString(buf, c.secretAccessKey)
	buf = appendLPString(buf, c.sessionToken)
	var exp [8]byte
	binary.BigEndian.PutUint64(exp[:], c.expiryUnixSecs)
	buf = append(buf, exp[:]...)
	if _, err := conn.Write(buf); err != nil {
		return fmt.Errorf("write ProvideCredentials: %w", err)
	}
	return readOkStatus(conn)
}

// appendLPString appends a [len:u16][bytes] field (big-endian length).
func appendLPString(buf []byte, s string) []byte {
	var l [2]byte
	binary.BigEndian.PutUint16(l[:], uint16(len(s)))
	buf = append(buf, l[:]...)
	return append(buf, s...)
}
