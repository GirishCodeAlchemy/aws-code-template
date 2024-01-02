package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler function for Lambda
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Print the event received
	eventJSON, _ := json.Marshal(request)
	fmt.Println("Received Event:", string(eventJSON))

	// Your Lambda logic goes here
	message := "Revcieved message: " + string(eventJSON)
	// Creating a response
	responseBody, err := json.Marshal(map[string]interface{}{
		"message": message,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(responseBody),
	}, nil
}

func main() {
	// Start the Lambda handler
	lambda.Start(handler)
}
