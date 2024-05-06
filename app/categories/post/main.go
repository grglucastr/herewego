package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
)

type Category struct {
	ID   string `json:"id" dynamodbav:"id"`
	Name string `json:"name" dynamodbav:"name"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Request Body")
	req := strings.ReplaceAll(request.Body, "\n", "")
	req = strings.ReplaceAll(req, "\r", "")

	fmt.Printf("%v", req)

	var cat Category
	json.Unmarshal([]byte(request.Body), &cat)

	cat.ID = uuid.NewString()
	r, _ := json.Marshal(cat)
	resp := string(r)

	AddCategory(cat)

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Body:       resp,
	}, nil
}

func main() {
	lambda.Start(handler)
}

func cloudConfigure() aws.Config {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))

	if err != nil {
		log.Fatalf("Failed to load configuration, %v", err)
	}

	return cfg

}

func AddCategory(cat Category) {

	category, err := attributevalue.MarshalMap(cat)

	if err != nil {
		panic(err)
	}

	cfg := cloudConfigure()
	client := dynamodb.NewFromConfig(cfg)

	_, err = client.PutItem(
		context.TODO(),
		&dynamodb.PutItemInput{
			TableName: aws.String("categories"),
			Item:      category,
		},
	)

	if err != nil {
		log.Fatalf("Couldn't add item to table. Here's why: %v\n", err)
	}
}
