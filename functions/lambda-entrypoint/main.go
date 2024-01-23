package main

import (
	"context"
	"github.com/magnuswahlstrand/sql-exercises/functions/app"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
)

var fiberLambda *fiberadapter.FiberLambda

// init the Fiber Server
func init() {
	log.Printf("Fiber cold start")
	fiberLambda = fiberadapter.New(app.SetupApp())
}

// Handler will deal with Fiber working with Lambda
func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return fiberLambda.ProxyWithContextV2(ctx, req)
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(Handler)
}
