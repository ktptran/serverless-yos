package main

import (
  "context"
  "fmt"

  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (e\
vents.APIGatewayProxyResponse, error) {
  fmt.Printf("Body size = %d.\n", len(request.Body))
  fmt.Println("Headers:")
  for key, value := range request.Headers {
    fmt.Printf("    %s: %s\n", key, value)
  }
  return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {
  lambda.Start(HandleRequest)
}
