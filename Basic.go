package restclient

import (
	"encoding/base64"
)

type Basic struct {
	username string
	password string
}

func NewBasic(username string, password string) Authorization {
	return Basic{
		username: username,
		password: password,
	}
}

func (b Basic) GetAuthorization() string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(b.username+":"+b.password))
}

func (b Basic) GetHeaderKey() string {
	return "Authorization"
}
