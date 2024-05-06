package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {

	req := events.APIGatewayProxyRequest{
		Body: "{'name':'George', 'age':30, 'city':'Curitiba'}",
	}

	res, _ := handler(req)

	expectedResponse := "Please return a list of categories"
	expectedStatusCode := 200

	if res.StatusCode != 200 {
		t.Errorf("Expected Status: %v, but got: %v", expectedStatusCode, res.StatusCode)
	}

	if res.Body != expectedResponse {
		t.Errorf("Expected Response: %v, but got: %v", expectedResponse, res.Body)
	}
}
