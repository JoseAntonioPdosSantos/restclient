package restclient

type HttpMethod string

const (
	Post   HttpMethod = "POST"
	Put    HttpMethod = "PUT"
	Get    HttpMethod = "GET"
	Delete HttpMethod = "DELETE"
)
