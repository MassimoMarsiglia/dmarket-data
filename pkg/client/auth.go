package client

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

var ErrFailedToDecode = errors.New("failed to decode")

func NewDmarketAuth(privKeyHex, pubKeyHex string) (*DmarketAuth, error) {
	privKey, err := hex.DecodeString(privKeyHex)
	if err != nil {
		return nil, fmt.Errorf("%w private key: %w", ErrFailedToDecode, err)
	}
	pubKey, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return nil, fmt.Errorf("%w public key: %w", ErrFailedToDecode, err)
	}
	return &DmarketAuth{
		privKey: privKey,
		pubKey:  pubKey,
	}, nil
}

type DmarketAuth struct {
	privKey []byte
	pubKey  []byte
}

func (d *DmarketAuth) BuildStringToSign(method, apiUrlPath, body, timestamp string) string {
	return fmt.Sprintf("%s%s%s%s", method, apiUrlPath, body, timestamp)
}

func (d *DmarketAuth) SignString(stringToSign string) (string, error) {
	signatureBytes := ed25519.Sign(d.privKey, []byte(stringToSign))
	return hex.EncodeToString(signatureBytes), nil
}

func (d *DmarketAuth) Middleware() func(ctx context.Context, req *http.Request) error {
	return func(ctx context.Context, req *http.Request) error {
		method := req.Method
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)
		apiUrlPath := req.URL.Path + req.URL.Query().Encode()
		var bodyBytes []byte
		var err error

		// 2. Read from req.Body instead of GetBody()
		if req.Body != nil && req.Body != http.NoBody {
			bodyBytes, err = io.ReadAll(req.Body)
			if err != nil {
				return err
			}
			// 3. IMPORTANT: Put the bytes back into the body so
			// the next handler can read it.
			req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		}

		stringToSign := d.BuildStringToSign(method, apiUrlPath, string(bodyBytes), timestamp)
		signature, err := d.SignString(stringToSign)

		req.Header.Set("X-Api-Key", string(d.pubKey))
		req.Header.Set("X-Request-Sign", signature)
		req.Header.Set("X-Request-Date", timestamp)
		return nil
	}
}
