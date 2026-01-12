package client

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func NewDmarketAuth(privKeyHex, pubKeyHex string) (*DmarketAuth, error) {
	privKey, err := hex.DecodeString(privKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key: %w", err)
	}
	pubKey, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode public key: %w", err)
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

func (d *DmarketAuth) BuildStringToSign(method, apiUrlPath, body string, timestamp string) string {
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
		reader, err := req.GetBody()
		if err != nil {
			return err
		}
		bodyBytes := make([]byte, 0)
		if reader != nil {
			bodyBytes, err = io.ReadAll(reader)
			if err != nil {
				return err
			}
		}
		stringToSign := d.BuildStringToSign(method, apiUrlPath, string(bodyBytes), timestamp)
		signature, err := d.SignString(stringToSign)

		req.Header.Set("X-Api-Key", string(d.pubKey))
		req.Header.Set("X-Request-Sign", signature)
		req.Header.Set("X-Request-Date", timestamp)
		return nil
	}
}
