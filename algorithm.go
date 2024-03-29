package restclient

type Algorithm interface {
	Prefix() string
	Name() string
	Exec(payload []byte) [32]byte
	Sign(doc []byte, secret []byte) []byte
}
