package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"tfnotifications/pkg/tfc"

	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"
)

func (h *handler) handle(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	ctx = context.WithValue(ctx, "log", logger)
	resp := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	if event.Body == "" {
		log.Fatal("body empty")
		resp.StatusCode = http.StatusBadRequest
	} else if event.HTTPMethod != "POST" {
		log.Fatal("wrong method: ", event.HTTPMethod)
		resp.StatusCode = http.StatusBadRequest
	}
	var tfc tfc.TerraformWebhook
	err := json.Unmarshal([]byte(event.Body), &tfc)
	if err != nil {
		log.Fatal("unmarshal request body failed",
			zap.Error(err))
		resp.StatusCode = http.StatusBadRequest
	}

	h.dclient.SendDiscordMessageFromTFC(ctx, tfc)
	return resp, nil
}
