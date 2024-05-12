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
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Category struct {
	ID   string `json:"id" dynamodbav:"id"`
	Name string `json:"name" dynamodbav:"name"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	id := request.PathParameters["id"]

	var newCategory Category
	json.Unmarshal([]byte(request.Body), &newCategory)
	newCategory.ID = id

	updatedValues := update(newCategory)

	resp, err := json.Marshal(updatedValues)

	if err != nil {
		log.Fatalln("Error while marshing the json: ", err)
		panic(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(resp),
	}, nil
}

func update(cat Category) Category {
	var err error
	var response *dynamodb.UpdateItemOutput
	var attributemap Category

	update := expression.Set(expression.Name("name"), expression.Value(cat.Name))

	expr, err := expression.NewBuilder().WithUpdate(update).Build()

	if err != nil {
		log.Fatalln("Failed to build update. Reason: ", err)
		panic(err)
	}

	client := dynamodb.NewFromConfig(getDbConfig())

	response, err = client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:                 aws.String("categories"),
		Key:                       getKey(cat.ID),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ReturnValues:              types.ReturnValueUpdatedNew,
	})

	if err != nil {
		log.Fatalln("Failed to update item. Reason: ", err)
		panic(err)
	}

	err = attributevalue.UnmarshalMap(response.Attributes, &attributemap)

	if err != nil {
		log.Fatalln("Failed to unmarsh the values. Reason: ", err)
		panic(err)
	}

	attributemap.ID = cat.ID
	return attributemap

}

func getDbConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))

	if err != nil {
		log.Fatalln("Failed to load config. Reason: ", err)
		panic(err)
	}

	return cfg
}

func getKey(id string) map[string]types.AttributeValue {

	catId, err := attributevalue.Marshal(id)

	if err != nil {
		log.Fatalln("Failed to generate key for update. Reason: ", err)
		panic(err)
	}

	return map[string]types.AttributeValue{"id": catId}

}

func main() {
	lambda.Start(handler)
}
