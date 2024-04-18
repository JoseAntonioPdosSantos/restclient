package restclient

type ContentType string

const ContentTypeDescription = "Content-Type"
const (
	ApplicationJson ContentType = "application/json"
	FormEncoded     ContentType = "application/x-www-form-urlencoded"
)
