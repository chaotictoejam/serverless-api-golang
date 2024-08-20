package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/rs/zerolog/log"
)

type Recipe struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
}

func getRecipeByID(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := req.PathParameters["id"]
	if id == "" {
		return ClientError(http.StatusBadRequest, "id is required")
	}

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to load config")
		return serverError(err)
	}

	client := dynamodb.NewFromConfig(cfg)

	output, err := client.GetItemRequest(&dynamodb.GetItemInput{
		TableName: aws.String("Recipes"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	}).Send(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get item")
		return serverError(err)
	}

	if output.Item == nil {
		return clientError(http.StatusNotFound, "recipe not found")
	}

	recipe := Recipe{}
	err = attributevalue.UnmarshalMap(output.Item, &recipe)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal recipe")
		return serverError(err)
	}

	body, err := json.Marshal(recipe)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal recipe")
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(getRecipeByID)
}
