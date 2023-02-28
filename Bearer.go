package restclient

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
	return "Bearer " + b.token
}
