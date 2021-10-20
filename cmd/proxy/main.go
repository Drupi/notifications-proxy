package main

import (
	"tfnotifications/pkg/discord"

	"github.com/aws/aws-lambda-go/lambda"
)

type handler struct {
	dclient discord.Discord
}

func main() {

	h := handler{
		dclient: discord.NewClient(),
	}
	lambda.Start(h.handle)
}
