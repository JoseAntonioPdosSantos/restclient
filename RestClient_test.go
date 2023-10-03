package restclient

import (
	"testing"
)

type ResponseBody struct {
	Id                 int      `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Price              int      `json:"price"`
	DiscountPercentage float64  `json:"discountPercentage"`
	Rating             float64  `json:"rating"`
	Stock              int      `json:"stock"`
	Brand              string   `json:"brand"`
	Category           string   `json:"category"`
	Thumbnail          string   `json:"thumbnail"`
	Images             []string `json:"images"`
}

func Test_WhenExecuteGetRequestWithUserUnauthorized_ThenReturnStatus401(t *testing.T) {
	authorization := NewBasic("your_username", "")
	httpClient := NewRestClient()
	httpResponse := httpClient.Delete("https://api.pagar.me/core/v5/customers/customer_id/cards/card_id").
		ContentType(ApplicationJson).
		Accept(ApplicationJson).
		Authorization(authorization).
		Exec()

	if httpResponse.GetError() != nil {
		t.Fatalf(`Expected value to be nil: %v`, httpResponse.GetError())
	}

	if httpResponse.GetResponse().StatusCode != 401 {
		t.Fatalf(`Expected value to be nil: %v`, httpResponse.GetResponse().StatusCode)
	}

}

func Test_WhenExecuteGetByUrl_ThenReturnResponseObject(t *testing.T) {

	httpClient := NewRestClient()
	httpResponse := httpClient.
		Get("https://dummyjson.com/products/30").
		ContentType(ApplicationJson).
		Accept(ApplicationJson).
		Exec()

	if httpResponse.GetError() != nil {
		t.Fatalf(`Expected value to be nil: %v`, httpResponse.GetError())
	}

	responseBody := &ResponseBody{}
	err := httpResponse.Unmarshal(&responseBody)

	if err != nil {
		t.Fatalf(`Expected value to be nil: %v`, err)
	}

	if responseBody.Id <= 0 || responseBody.Title == "" {
		t.Fatalf(`Expected value to be nil: %v`, err)
	}

}

func Test_WhenExecuteGet_ThenReturnResponseObject(t *testing.T) {

	httpClient := NewRestClient()
	httpResponse := httpClient.
		Get("https://dummyjson.com/products/30").
		ContentType(ApplicationJson).
		Accept(ApplicationJson).
		Exec()

	if httpResponse.GetError() != nil {
		t.Fatalf(`Expected value to be nil: %v`, httpResponse.GetError())
	}

	responseBody := &ResponseBody{}
	err := httpResponse.Unmarshal(&responseBody)

	if err != nil {
		t.Fatalf(`Expected value to be nil: %v`, err)
	}

	if responseBody.Id <= 0 || responseBody.Title == "" {
		t.Fatalf(`Expected value to be nil: %v`, err)
	}

}
