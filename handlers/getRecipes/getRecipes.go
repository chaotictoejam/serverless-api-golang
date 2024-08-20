package main

import (
	"context"
	"encoding/json"
	"net/http"
	"serverless-api-go-example/helpers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/rs/zerolog/log"
)

type Recipe struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
}

func GetRecipes(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to load config")
		return helpers.ServerError(err)
	}

	client := dynamodb.NewFromConfig(cfg)

	output, err := client.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String("Recipes"),
	})

	if err != nil {
		log.Error().Err(err).Msg("failed to scan table")
		return helpers.ServerError(err)
	}

	recipes := make([]Recipe, 0, len(output.Items))
	for _, item := range output.Items {
		recipe := Recipe{}
		err = attributevalue.UnmarshalMap(item, &recipe)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal recipe")
			return helpers.ServerError(err)
		}
		recipes = append(recipes, recipe)
	}

	body, err := json.Marshal(recipes)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal recipes")
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
	lambda.Start(GetRecipes)
}
