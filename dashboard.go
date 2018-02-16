package gofaas

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

// Dashboard returns a dashboard HTML page
func Dashboard(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body: string("<html><body><h1>GoFAAS Dashboard</h1></body></html>"),
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
		StatusCode: 200,
	}, nil
}
