package restclient

type Authorization interface {
	GetAuthorization() string
	GetHeaderKey() string
}
