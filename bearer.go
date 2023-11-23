package restclient

import "fmt"

type Bearer struct {
	token  string
	prefix string
}

func NewBearer(token string) Authorization {
	return Bearer{
		token:  token,
		prefix: "Bearer",
	}
}
func (b Bearer) GetAuthorization() string {
	return fmt.Sprintf("%s %s", b.prefix, b.token)
}

func (b Bearer) GetHeaderKey() string {
	return "Authorization"
}
