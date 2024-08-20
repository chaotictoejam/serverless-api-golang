package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type errorBody struct {
	Error string `json:"error"`
}

func ServerError(err error) (events.APIGatewayProxyResponse, error) {
	body, _ := json.Marshal(errorBody{
		Error: "Internal Server Error",
	})

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func ClientError(status int, message ...string) (events.APIGatewayProxyResponse, error) {
	body, _ := json.Marshal(errorBody{
		Error: fmt.Sprint(message),
	})

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
