package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/morph-l2/externalsign"
	"github.com/morph-l2/go-ethereum/log"
)

// genAddressPath is the morph-proxy endpoint that asks the sign machine to
// create a new wallet address.
const genAddressPath = "/v1/sign/gen_address"

// ExternalAddrGenerator requests new external sign wallet addresses from
// morph-proxy, using the same app credentials as transaction signing.
type ExternalAddrGenerator struct {
	es  *externalsign.ExternalSign
	url string
}

// NewExternalAddrGenerator derives the gen_address url from cfg.Url (the
// tx_sign url) by replacing its path.
func NewExternalAddrGenerator(cfg ExternalSignConfig) (*ExternalAddrGenerator, error) {
	u, err := url.Parse(cfg.Url)
	if err != nil {
		return nil, fmt.Errorf("parse external sign url: %w", err)
	}
	u.Path = genAddressPath
	u.RawQuery = ""
	return &ExternalAddrGenerator{
		es:  externalsign.NewExternalSign(cfg.Appid, cfg.RsaPriv, "", cfg.Chain, nil),
		url: u.String(),
	}, nil
}

// ServeHTTP handles POST requests: generates one new address, logs it and
// returns it as {"address": "0x..."}.
func (g *ExternalAddrGenerator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	addr, err := g.es.RequestWalletAddr(g.url)
	if err != nil {
		log.Error("generate external sign address failed", "url", g.url, "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	log.Info("generated external sign address", "address", addr.Hex(), "chain", g.es.Chain, "appid", g.es.Appid)
	_ = json.NewEncoder(w).Encode(map[string]string{"address": addr.Hex()})
}
