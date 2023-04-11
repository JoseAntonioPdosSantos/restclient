package restclient

import (
	"crypto/sha256"
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

func (s SHA256) Exec(payload []byte) [32]byte {
	return sha256.Sum256(payload)
}
