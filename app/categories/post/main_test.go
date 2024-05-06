package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {

	req := events.APIGatewayProxyRequest{
		Body: "{\"categoria\":\"Lazer\"}",
	}

	res, err := handler(req)

	if err != nil {
		t.Errorf("Unexpected error")
	}

	expectedRes := "work in progress"
	expectedStatusCode := 201

	if res.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got: %v", expectedStatusCode, res.StatusCode)
	}

	if res.Body != expectedRes {
		t.Errorf("Expected response %v, but got: %v", expectedRes, res.Body)
	}

}
