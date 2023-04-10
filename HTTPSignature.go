package restclient

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
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
	headers          map[string]string
	header           []string
	canonicalHeaders *bytes.Buffer
	signature        string
	algorithm        Algorithm
}

func NewHTTPSignatureBuilder() hTTPSignatureBuilder {
	return hTTPSignatureBuilder{
		headers:          make(map[string]string),
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
	h.headers[key] = value
	h.header = append(h.header, key)
	fmt.Fprintf(h.canonicalHeaders, "%s: %s\n", key, value)
	return h
}

func (h hTTPSignatureBuilder) Host(host string) hTTPSignatureBuilder {
	h.headers["host"] = host
	h.header = append(h.header, "host")
	fmt.Fprintf(h.canonicalHeaders, "%s: %s\n", "host", host)
	return h
}

func (h hTTPSignatureBuilder) Date(date string) hTTPSignatureBuilder {
	h.headers["date"] = date
	h.header = append(h.header, "date")
	h.date = date
	fmt.Fprintf(h.canonicalHeaders, "%s: %s\n", "date", date)
	return h
}

func (h hTTPSignatureBuilder) RequestTarget(requestTarget string) hTTPSignatureBuilder {
	h.headers["(request-target)"] = requestTarget
	h.header = append(h.header, "(request-target)")
	fmt.Fprintf(h.canonicalHeaders, "%s: %s\n", "(request-target)", requestTarget)
	return h
}

func (h hTTPSignatureBuilder) Digest(payload []byte) hTTPSignatureBuilder {
	if h.algorithm == nil {
		return h
	}

	x := NewSHA256().Exec(payload)

	bodyReq := x.([]byte) //h.algorithmFn(payload)
	h.digest = fmt.Sprintf("%s=%s", h.algorithm.Prefix(), base64.StdEncoding.EncodeToString(bodyReq[:]))
	h.headers["digest"] = h.digest
	h.header = append(h.header, "digest")
	fmt.Fprintf(h.canonicalHeaders, "%s: %s\n", "digest", h.digest)
	return h
}

func (h hTTPSignatureBuilder) VCMerchantID(vCMerchantID string) hTTPSignatureBuilder {
	h.headers["v-c-merchant-id"] = vCMerchantID
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
	secretKey := hmac.New(sha256.New, decodedSecretKey)
	secretKey.Write([]byte(message))
	hash := secretKey.Sum(nil)

	hash_ := base64.StdEncoding.EncodeToString(hash)

	return hash_
}

func (s HTTPSignature) GetAuthorization() string {
	return s.Signature
}

func (s HTTPSignature) GetHeaderKey() string {
	return "signature"
}
