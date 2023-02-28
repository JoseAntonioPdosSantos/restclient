package restclient

type RestClient struct {
}

func NewRestClient() HttpClientMethod {
	return RestClient{}
}

func (r RestClient) Get() HttpIntegration {
	return r.newHTTPMethod(Get)
}

func (r RestClient) Post() HttpIntegration {
	return r.newHTTPMethod(Post)
}

func (r RestClient) Delete() HttpIntegration {
	return r.newHTTPMethod(Delete)
}

func (r RestClient) newHTTPMethod(method HttpMethod) HttpIntegration {
	return HTTPClient{
		httpMethod:    method,
		url:           "https://",
		header:        make(map[string]string),
		params:        make(map[string]string),
		queries:       make(map[string]string),
		authorization: nil,
	}
}
