# Introduction 
The restclient library is an HTTP client designed to make HTTP requests easier.

# Getting Started
This library provides several HTTP request solutions, as demonstrated in the examples below.

> If you can't find what you're looking for, you can help us implement the solution by sending 
> a pull request or letting us know. We're open to implementing it for you. 
> Take a look at the [Contribute Section](https://github.com/JoseAntonioPdosSantos/restclient/blob/master/contributing.md).

To import this library into your Golang project, use the following command:
```cmd
go get github.com/JoseAntonioPdosSantos/restclient
```

### The solutions implemented in this library are described in the examples below:

Before making a request, you need to create an httpClient object:
``` go
httpClient := restclient.NewRestClient()
```
Make sure to import the library correctly:
```go
import (
	"github.com/JoseAntonioPdosSantos/restclient"
)
```

#### Making a simple HTTP request

```go
response := httpClient.Get().
	Url("https://viacep.com.br/ws/01001000/json/").
	Exec()

body, _ := response.GetBody()

fmt.Printf("data: %v", string(body))

```
#### Making a simple HTTP request with parameters

```go

response := httpClient.Get().
	Url("https://viacep.com.br/ws/${cep}/json/").
	AddParams("cep", "01001000").
	Exec()

body, _ := response.GetBody()

fmt.Printf("data: %v", string(body))

```

#### Using basic authentication in your request

```go
response := httpClient.Get().
	Url("https://your-rest-api-integration-herer").
	Authorization(restclient.NewBasic("your username", "your password")).
	Exec()
```


#### Using HTTP Signature authentication with SHA256 in your POST request

```GO
httpSignatureAuthorization := restclient.NewHTTPSignatureBuilder().
	Algorithm(restclient.NewSHA256()).
	KeyID("Your_Key_ID").
	SharedSecretKey("Your_Shared_Secret_Key").
	Host("HOST").
	Date(time.Now().UTC().Format(time.RFC1123)).
	RequestTarget(fmt.Sprintf(
		"%s %s", 
		strings.ToLower(string(restclient.Post)), 
		"HOST",
		).
	Digest(body).
	VCMerchantID("Your_Merchant_ID").
	Build()

response := httpClient.Post().
	Url(url).
	ContentType(restclient.ApplicationJson).
	AddHeader("Digest", authorization.Digest).
	AddHeader("V-C-Merchant-Id", "Your_Merchant_ID").
	AddHeader("Date", authorization.Date).
	AddHeader("Host", "Your_Host").
	AddParams("id", paymentId).
	Authorization(httpSignatureAuthorization).
	Body(body).
	Exec()
	
```

#### Using other configurations in your request

```go
response := httpClient.Get().
	Url("https://your-rest-api-integration-herer").
	ContentType(restclient.ApplicationJson).
	Accept(restclient.ApplicationJson).
	AddHeader("your_key", "your_value").
	AddParams("your_key", "your_value").
	Authorization(restclient.NewBasic("your username", "your password")).
	Interceptor(your_Interceptor_Implemented_Here).
	Body(body).
	Exec()
```

#### Using an interceptor in your request
First, you need to implement the `http.RoundTripper` interface from the `net/http` package in GoLang, 
and then pass your implementation as shown below:

```go
response := httpClient.Get().
	Url("https://your-rest-api-integration-herer").
	Interceptor(your_Interceptor_Implemented_Here).
	Exec()
```
