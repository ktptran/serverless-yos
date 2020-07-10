package main

import (
  "context"
  "log"
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/aws/aws-lambda-go/lambdacontext"
)

// lc is used to consume the information that the context object caputred
func Handler(ctx context.Context) {
  lc, _ := lambdacontext.FromContext(ctx)
  log.print(lc.AwsRequestId)
}

func main() {
  lambda.Start(Handler)
}
