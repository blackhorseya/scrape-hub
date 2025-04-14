package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(c context.Context) (string, error) {
	return "ok", nil
}

func main() {
	lambda.Start(Handler)
}
