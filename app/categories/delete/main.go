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

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	id := request.PathParameters["id"]

	delete(id)

	return events.APIGatewayProxyResponse{
		Body:       "{}",
		StatusCode: 204,
	}, nil
}

func getDbConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))

	if err != nil {
		log.Fatalln("Failed to load config. Reason: ", err)
	}

	return cfg
}

func getKey(id string) map[string]types.AttributeValue {

	catId, err := attributevalue.Marshal(id)

	if err != nil {
		log.Fatalln("Failed do get the key. Reason: ", catId)
	}

	return map[string]types.AttributeValue{"id": catId}

}

func delete(id string) {

	client := dynamodb.NewFromConfig(getDbConfig())

	params := &dynamodb.DeleteItemInput{
		TableName: aws.String("categories"),
		Key:       getKey(id),
	}

	_, err := client.DeleteItem(context.TODO(), params)

	if err != nil {
		log.Fatalln("Failed to delete. Reason: ", err)
		panic(err)
	}

}

func main() {
	lambda.Start(handler)
}
