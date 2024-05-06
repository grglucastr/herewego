package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Category struct {
	ID   string `json:"id" dynamodbav:"id"`
	Name string `json:"name" dynamodbav:"name"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	categories, err := listAllDB()

	if err != nil {
		panic(err)
	}

	json, _ := json.Marshal(categories)
	resBody := string(json)
	log.Printf("Response: %v\n", resBody)

	return events.APIGatewayProxyResponse{
		Body:       resBody,
		StatusCode: 200,
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

func listAllDB() ([]Category, error) {

	client := dynamodb.NewFromConfig(cloudConfigure())

	params := &dynamodb.ScanInput{
		TableName: aws.String("categories"),
	}
	result, err := client.Scan(context.TODO(), params)

	if err != nil {
		log.Fatalf("Query API failed: %s\n", err)
	}

	var categories []Category
	for _, item := range result.Items {

		cat := Category{}

		err = attributevalue.UnmarshalMap(item, &cat)

		if err != nil {
			log.Fatalf("Got error unmarshalling: %s", err)
		}

		categories = append(categories, cat)
	}

	return categories, nil
}
