package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// BeaconClient wraps Beacon Chain API client
type BeaconClient struct {
	rpcURL     string
	httpClient *http.Client
}

// NewBeaconClient creates new Beacon client
func NewBeaconClient(rpcURL string) *BeaconClient {
	return &BeaconClient{
		rpcURL: strings.TrimSuffix(rpcURL, "/"),
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// BlobSidecar represents beacon chain blob sidecar structure
type BlobSidecar struct {
	Index             string `json:"index"`
	Blob              string `json:"blob"`
	KZGCommitment     string `json:"kzg_commitment"`
	KZGProof          string `json:"kzg_proof"`
	SignedBlockHeader struct {
		Message struct {
			Slot          string `json:"slot"`
			ProposerIndex string `json:"proposer_index"`
		} `json:"message"`
	} `json:"signed_block_header"`
}

// BlobSidecarsResponse represents beacon API response
type BlobSidecarsResponse struct {
	Data []BlobSidecar `json:"data"`
}

// QueryBlobSidecars queries blob sidecars by block root
// blockRoot: block root hash (hex string with 0x prefix)
// indices: blob index list (optional, nil for all blobs)
func (c *BeaconClient) QueryBlobSidecars(ctx context.Context, blockRoot string, indices []uint64) (*BlobSidecarsResponse, error) {
	return c.queryBlobSidecars(ctx, blockRoot, indices)
}

// QueryBlobSidecarsBySlot queries blob sidecars by slot number
// slot: beacon chain slot number
// indices: blob index list (optional, nil for all blobs)
func (c *BeaconClient) QueryBlobSidecarsBySlot(ctx context.Context, slot uint64, indices []uint64) (*BlobSidecarsResponse, error) {
	blockID := fmt.Sprintf("%d", slot)
	return c.queryBlobSidecars(ctx, blockID, indices)
}

// QueryBlobSidecarsByBlockID queries blob sidecars by block identifier
// blockID can be:
// - "head": latest block
// - "finalized": latest finalized block
// - slot number (e.g. "12345")
// - block root (hex string with 0x prefix)
func (c *BeaconClient) QueryBlobSidecarsByBlockID(ctx context.Context, blockID string, indices []uint64) (*BlobSidecarsResponse, error) {
	return c.queryBlobSidecars(ctx, blockID, indices)
}

// queryBlobSidecars internal method to query blob sidecars
func (c *BeaconClient) queryBlobSidecars(ctx context.Context, blockID string, indices []uint64) (*BlobSidecarsResponse, error) {
	// Build URL
	url := fmt.Sprintf("%s/eth/v1/beacon/blob_sidecars/%s", c.rpcURL, blockID)

	// Add index parameters
	if len(indices) > 0 {
		var indexStrs []string
		for _, idx := range indices {
			indexStrs = append(indexStrs, fmt.Sprintf("%d", idx))
		}
		url = fmt.Sprintf("%s?indices=%s", url, strings.Join(indexStrs, ","))
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to query blob sidecars: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("beacon API returned status %d: %s", resp.StatusCode, string(body))
	}

	// Parse response
	var result BlobSidecarsResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// GetBeaconBlockHeader gets beacon block header by block ID
func (c *BeaconClient) GetBeaconBlockHeader(ctx context.Context, blockID string) (*BeaconBlockHeader, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/headers/%s", c.rpcURL, blockID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to query block header: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("beacon API returned status %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Data BeaconBlockHeader `json:"data"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result.Data, nil
}

// BeaconBlockHeader represents a beacon block header
type BeaconBlockHeader struct {
	Root      string `json:"root"`
	Canonical bool   `json:"canonical"`
	Header    struct {
		Message struct {
			Slot          string `json:"slot"`
			ProposerIndex string `json:"proposer_index"`
			ParentRoot    string `json:"parent_root"`
			StateRoot     string `json:"state_root"`
			BodyRoot      string `json:"body_root"`
		} `json:"message"`
		Signature string `json:"signature"`
	} `json:"header"`
}
