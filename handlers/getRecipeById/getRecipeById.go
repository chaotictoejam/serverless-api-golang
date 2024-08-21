package main

import (
	"context"
	"encoding/json"
	"net/http"
	"serverless-api-go-example/helpers"
	"serverless-api-go-example/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/rs/zerolog/log"
)

func getRecipeByID(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := req.PathParameters["id"]
	if id == "" {
		return helpers.ClientError(http.StatusBadRequest, "id is required")
	}

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to load config")
		return helpers.ServerError(err)
	}

	client := dynamodb.NewFromConfig(cfg)

	output, err := client.GetItem(ctx,
		&dynamodb.GetItemInput{
			TableName: aws.String("Recipes"),
			Key: map[string]types.AttributeValue{
				"id": &types.AttributeValueMemberS{Value: id},
			},
		})

	if err != nil {
		log.Error().Err(err).Msg("failed to get item")
		return helpers.ServerError(err)
	}

	if output.Item == nil {
		return helpers.ClientError(http.StatusNotFound, "recipe not found")
	}

	recipe := models.Recipe{}
	err = attributevalue.UnmarshalMap(output.Item, &recipe)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal recipe")
		return helpers.ServerError(err)
	}

	body, err := json.Marshal(recipe)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal recipe")
		return helpers.ServerError(err)
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
