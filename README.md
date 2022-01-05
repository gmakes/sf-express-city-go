# sf-express-city-go
顺丰同城开放平台 golang sdk

## 官方文档

https://commit-openic.sf-express.com/open/api/docs/index#/apidoc

## 使用方法

```go
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

	t.Logf("resp: %v", resp)
}
```