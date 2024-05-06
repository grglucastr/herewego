package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Category struct {
	ID   string `json:"id" dynamodbav:"id"`
	Name string `json:"name" dynamodbav:"name"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	params := request.PathParameters
	id := params["id"]

	log.Println("Serching category for the following id: {}", id)

	loadGetById(id)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "{}",
	}, nil

}

func main() {
	lambda.Start(handler)
}

func getDbConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))

	if err != nil {
		log.Fatalf("Failed to load configuration, %v", err)
	}

	return cfg
}

func loadGetById(id string) Category {

	cfg := getDbConfig()
	client := dynamodb.NewFromConfig(cfg)

	log.Println("The key: {}", getKey(id))

	params := &dynamodb.GetItemInput{
		Key: getKey(id), TableName: aws.String("categories"),
	}

	resp, err := client.GetItem(context.TODO(), params)

	if err != nil {
		panic(err)
	}

	var cat Category
	err = attributevalue.UnmarshalMap(resp.Item, &cat)

	if err != nil {
		panic(err)
	}

	return cat
}

func getKey(id string) map[string]types.AttributeValue {
	catId, err := attributevalue.Marshal(id)

	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"id": catId}
}
