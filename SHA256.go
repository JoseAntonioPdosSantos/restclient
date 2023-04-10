package restclient

import (
	"crypto/sha256"
	"encoding/json"
)

type SHA256 struct {
}

func NewSHA256() Algorithm {
	return SHA256{}
}

func (s SHA256) Prefix() string {
	return "SHA-256"
}

func (s SHA256) Name() string {
	return "HmacSHA256"
}

func (s SHA256) Exec(payload interface{}) interface{} {
	payload_, _ := json.Marshal(payload)
	return sha256.Sum256(payload_)
}
