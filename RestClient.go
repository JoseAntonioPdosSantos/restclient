package restclient

type RestClient struct {
}

func NewRestClient() HttpClientMethod {
	return RestClient{}
}

func (r RestClient) Post() HttpIntegration {
	return r.newHTTPMethod(Post)
}

func (r RestClient) Put() HttpIntegration {
	return r.newHTTPMethod(Put)
}

func (r RestClient) Get() HttpIntegration {
	return r.newHTTPMethod(Get)
}

func (r RestClient) Delete() HttpIntegration {
	return r.newHTTPMethod(Delete)
}

func (r RestClient) newHTTPMethod(method HttpMethod) HttpIntegration {
	return HttpClient{
		httpMethod:    method,
		url:           "https://",
		header:        make(map[string]string),
		params:        make(map[string]string),
		queries:       make(map[string]string),
		authorization: nil,
	}
}
