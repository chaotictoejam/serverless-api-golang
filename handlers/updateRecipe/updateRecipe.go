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
	"github.com/rs/zerolog/log"
)

func UpdateRecipe(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to load config")
		return helpers.ServerError(err)
	}

	client := dynamodb.NewFromConfig(cfg)

	var recipe models.Recipe
	err = json.Unmarshal([]byte(req.Body), &recipe)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal recipe")
		return helpers.ServerError(err)
	}

	av, err := attributevalue.MarshalMap(recipe)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal recipe")
		return helpers.ServerError(err)
	}

	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("Recipes"),
		Item:      av,
	})

	if err != nil {
		log.Error().Err(err).Msg("failed to put item")
		return helpers.ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Recipe updated successfully",
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
	}, nil
}

func main() {
	lambda.Start(UpdateRecipe)
}
