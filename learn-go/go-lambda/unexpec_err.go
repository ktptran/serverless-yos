package main

import (
  "errors"

  "github.com/aws/aws-lambda-go/lambda"
)

func handler(string) (string, error) {
  panic(errors.New("Something went wrong"))
}

func main() {
  lambda.Start(handler)
}
