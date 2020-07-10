package main

import (
  "fmt"
  "os"
  "github.com/aws/aws-lambda-go/lambda"
)

func main() {
  fmt.Printf("Lambda is in the %s region and is on %s", os.Getenv("AWS_REGION"),
    os.Getenv("AWS_EXECUTION_ENV"))
}
