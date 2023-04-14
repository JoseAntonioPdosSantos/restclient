package restclient

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
)

type HTTPSignature struct {
	Signature string
	Digest    string
	Date      string
}

type hTTPSignatureBuilder struct {
	keyID            string
	sharedSecretKey  string
	merchantID       string
	date             string
	digest           string
	header           []string
	canonicalHeaders *bytes.Buffer
	signature        string
	algorithm        Algorithm
}

func NewHTTPSignatureBuilder() hTTPSignatureBuilder {
	return hTTPSignatureBuilder{
		header:           make([]string, 0, 10),
		canonicalHeaders: bytes.NewBuffer([]byte{}),
	}
}

func (h hTTPSignatureBuilder) Algorithm(algorithm Algorithm) hTTPSignatureBuilder {
	h.algorithm = algorithm
	return h
}

func (h hTTPSignatureBuilder) KeyID(keyID string) hTTPSignatureBuilder {
	h.keyID = keyID
	return h
}

func (h hTTPSignatureBuilder) SharedSecretKey(sharedSecretKey string) hTTPSignatureBuilder {
	h.sharedSecretKey = sharedSecretKey
	return h
}

func (h hTTPSignatureBuilder) MerchantID(merchantID string) hTTPSignatureBuilder {
	h.merchantID = merchantID
	return h
}

func (h hTTPSignatureBuilder) AddHeader(key string, value string) hTTPSignatureBuilder {
	h.header = append(h.header, key)
	fmt.Fprintf(h.canonicalHeaders, "%s: %s\n", key, value)
	return h
}

func (h hTTPSignatureBuilder) Host(host string) hTTPSignatureBuilder {
	h.header = append(h.header, "host")
	fmt.Fprintf(h.canonicalHeaders, "%s: %s\n", "host", host)
	return h
}

func (h hTTPSignatureBuilder) Date(date string) hTTPSignatureBuilder {
	h.header = append(h.header, "date")
	h.date = date
	fmt.Fprintf(h.canonicalHeaders, "%s: %s\n", "date", date)
	return h
}

func (h hTTPSignatureBuilder) RequestTarget(requestTarget string) hTTPSignatureBuilder {
	h.header = append(h.header, "(request-target)")
	fmt.Fprintf(h.canonicalHeaders, "%s: %s\n", "(request-target)", requestTarget)
	return h
}

func (h hTTPSignatureBuilder) Digest(payload []byte) hTTPSignatureBuilder {
	if h.algorithm == nil {
		return h
	}
	bodyReq := h.algorithm.Exec(payload)
	h.digest = fmt.Sprintf("%s=%s", h.algorithm.Prefix(), base64.StdEncoding.EncodeToString(bodyReq[:]))
	h.header = append(h.header, "digest")
	fmt.Fprintf(h.canonicalHeaders, "%s: %s\n", "digest", h.digest)
	return h
}

func (h hTTPSignatureBuilder) VCMerchantID(vCMerchantID string) hTTPSignatureBuilder {
	h.header = append(h.header, "v-c-merchant-id")
	fmt.Fprintf(h.canonicalHeaders, "%s: %s\n", "v-c-merchant-id", vCMerchantID)
	return h
}

func (h hTTPSignatureBuilder) Build() HTTPSignature {
	canonicalString := strings.TrimSuffix(h.canonicalHeaders.String(), "\n")
	sign := h.sign(canonicalString, h.sharedSecretKey)

	signature := fmt.Sprintf(`keyid="%s", algorithm="%s", headers="%s", signature="%s"`,
		h.keyID,
		h.algorithm.Name(),
		strings.Join(h.header, " "),
		sign,
	)

	return HTTPSignature{
		Signature: signature,
		Digest:    h.digest,
		Date:      h.date,
	}
}

func (h hTTPSignatureBuilder) sign(message string, secret string) string {
	decodedSecretKey, _ := base64.StdEncoding.DecodeString(secret)
	hash := h.algorithm.Sign([]byte(message), decodedSecretKey)
	return base64.StdEncoding.EncodeToString(hash)
}

func (s HTTPSignature) GetAuthorization() string {
	return s.Signature
}

func (s HTTPSignature) GetHeaderKey() string {
	return "signature"
}
