package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {

	req := events.APIGatewayProxyRequest{
		Body: "{}",
	}

	res, err := handler(req)

	if err != nil {
		panic(err)
	}

	if res.StatusCode != 204 {
		t.Errorf("Expected: %v, but got: %v", 204, res.StatusCode)
	}

}
