package restclient

import (
	"testing"
)

func TestRestClientGet(t *testing.T) {
	authorization := NewBasic("your_username", "")
	httpClient := NewRestClient()
	_, err := httpClient.Delete().
		Url("https://api.pagar.me/core/v5/customers/customer_id/cards/card_id").
		ContentType(ApplicationJson).
		Accept(ApplicationJson).
		Authorization(authorization).
		Exec()

	if err != nil {
		t.Fatalf(`Expected value to be nil: %v`, err)
	}

}
