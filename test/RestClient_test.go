package test

import (
	restclient2 "rest-client-api/infra/restclient"
	"testing"
)

func TestRestClientGet(t *testing.T) {
	authorization := restclient2.NewBasic("your_username", "")
	httpClient := restclient2.NewRestClient()
	_, err := httpClient.Get().
		Url("https://api.pagar.me/core/v5/customers/customer_id/cards/card_id").
		AddHeader("accept", "application/json").
		AddHeader("content-type", "application/json").
		Authorization(authorization).
		Exec()

	if err != nil {
		t.Fatalf(`Expected value to be nil: %v`, err)
	}

}
