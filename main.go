package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type MyEvent struct {
	// sample properties below
	InstanceId string `json:"InstanceId"`
	Status     string `json:"Status"`
}

func HandleRequest(context context.Context, ev MyEvent) (events.APIGatewayProxyResponse, error) {

	res := events.APIGatewayProxyResponse{
		IsBase64Encoded: false,
		StatusCode:      200,
		Headers:         nil,
		Body:            "Success!",
	}

	region := os.Getenv("AWS_REGION")

	// Load session from shared config
	_, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		res.Body = "Error!"
	}

	return res, err
}

// entry point to your lambda
func main() {
	lambda.Start(HandleRequest)
}
