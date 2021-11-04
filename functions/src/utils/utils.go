package utils

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func CreateApiResponse(v interface{}, statusCode int) *events.APIGatewayProxyResponse {
	responseBody := ""
	if v != "" {
		jsonBody, _ := json.Marshal(v)
		responseBody = string(jsonBody)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "Content-Type",
			"Access-Control-Allow-Methods": "GET",
		},
		Body: responseBody,
	}
}

func ParseBody(body string, v interface{}) error {
	bytes := []byte(body)
	return json.Unmarshal(bytes, v)
}
