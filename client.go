package sfexpresscitygo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SFExpressCityClient struct {
	signer *Signer
}

func NewSFExpressCityClient(param *ClientParams) (*SFExpressCityClient, error) {
	return &SFExpressCityClient{
		signer: NewSigner(param.DevId, param.DevKey),
	}, nil
}

func (c *SFExpressCityClient) PreCreateOrder(request *PreCreateOrderRequest) (int, *Response, error) {
	code, resp, err := c.do(PRE_CREATE_ORDER_FORMAT, request)
	if err != nil {
		return code, resp, err
	}
	// TODO format response
	return code, resp, nil
}

func (c *SFExpressCityClient) do(apiFormat string, request interface{}) (int, *Response, error) {

	payload, err := json.Marshal(request)
	if err != nil {
		return 400, nil, err
	}

	signature := c.signer.Sign(string(payload))
	responseBody := bytes.NewBuffer([]byte(payload))

	url := fmt.Sprintf(apiFormat, signature)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(url, "application/json", responseBody)
	//Handle Error
	if err != nil {
		return 400, nil, err
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 400, nil, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 500, nil, err
	}

	return resp.StatusCode, &response, nil
}
