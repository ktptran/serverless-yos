package main

// Log files replace the fmt print statements

import (
  "log"
  "github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest() {
  log.print("Hello from Lambda")
}

func main() {
  lambda.Start(HandleRequest)
}
