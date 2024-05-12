package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {

	params := map[string]string{
		"id": "1111",
	}

	request := events.APIGatewayProxyRequest{
		Body:           "{}",
		PathParameters: params,
	}

	response, err := handler(request)

	if err != nil {
		t.Errorf("Something wrong happened")
	}

	if response.StatusCode != 200 {
		t.Errorf("Expected %v, but got: %v", 200, response.StatusCode)
	}

}
