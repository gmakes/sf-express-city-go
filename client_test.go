package sfexpresscitygo_test

import (
	"fmt"
	"testing"

	sfexpresscitygo "github.com/gmakes/sf-express-city-go"
)

func TestPreCreateOrder(t *testing.T) {
	expected := 200

	client, err := sfexpresscitygo.NewSFExpressCityClient(&sfexpresscitygo.ClientParams{
		DevId:  "",
		DevKey: "",
	})

	if err != nil {
		t.Error("Can not init client")
	}

	code, resp, err := client.PreCreateOrder(&sfexpresscitygo.PreCreateOrderRequest{
		DevId: 12345678,
		// ...
	})
	if err != nil {
		t.Error(err.Error())
	}

	actual := code
	if actual != expected {
		t.Errorf("PreCreateOrder status actual = %d; expected %d", actual, expected)
	}

	fmt.Printf("resp: %v\n", resp)
}
