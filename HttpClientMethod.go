package restclient

type HttpClientMethod interface {
	Post() HttpIntegration
	Put() HttpIntegration
	Get() HttpIntegration
	Delete() HttpIntegration
}
