package restclient

type RestClient struct {
}

func NewRestClient() HttpClientMethod {
	return RestClient{}
}

func (r RestClient) Post(url ...string) HttpIntegration {
	return r.newHTTPMethod(Post, url)
}

func (r RestClient) Put(url ...string) HttpIntegration {
	return r.newHTTPMethod(Put, url)
}

func (r RestClient) Get(url ...string) HttpIntegration {
	return r.newHTTPMethod(Get, url)
}

func (r RestClient) Delete(url ...string) HttpIntegration {
	return r.newHTTPMethod(Delete, url)
}

func (r RestClient) newHTTPMethod(method HttpMethod, url []string) HttpIntegration {
	client := HttpClient{
		httpMethod:    method,
		url:           "https://",
		header:        make(map[string]string),
		params:        make(map[string]string),
		queries:       make(map[string]string),
		authorization: nil,
	}
	if len(url) > 0 && url[0] != "" {
		client.url = url[0]
	}
	return &client
}
