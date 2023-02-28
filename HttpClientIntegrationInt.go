package restclient

type HttpClientMethod interface {
	Get() HttpIntegration
	Post() HttpIntegration
	Delete() HttpIntegration
}
