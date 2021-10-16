package main

import (
	"net/http"
	"os"
	"tfnotifications/pkg/discord"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	h := &handler{
		dclient: discord.Discord{
			Token: os.Getenv("DISCORD_API"),
			HTTPClient: &http.Client{
				Timeout: 30 * time.Second,
			},
		},
	}
	lambda.Start(h.handleRequest)
}
