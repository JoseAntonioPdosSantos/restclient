package restclient

type HttpClientMethod interface {
	// Post executes a single HTTP POST transaction, returning
	//
	// url is the url, host, or path of the request
	Post(url ...string) HttpIntegration
	// Put executes a single HTTP POST transaction, returning
	//
	// url is the url, host, or path of the request
	Put(url ...string) HttpIntegration
	// Get executes a single HTTP POST transaction, returning
	//
	// url is the url, host, or path of the request
	Get(url ...string) HttpIntegration
	// Delete executes a single HTTP POST transaction, returning
	//
	// url is the url, host, or path of the request
	Delete(url ...string) HttpIntegration
}
