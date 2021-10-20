package discord

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"tfnotifications/pkg/tfc"
	"time"

	"go.uber.org/zap"
)

type Discord struct {
	Token      string
	HTTPClient *http.Client
}

// DiscordWebhook is a partial struct for https://discord.com/developers/docs/resources/webhook#execute-webhook
type DiscordWebhook struct {
	Content string         `json:"content"`
	Embeds  []DiscordEmbed `json:"embeds"`
}

type DiscordEmbed struct {
	Title       string              `json:"title"`
	Description string              `json:"description"`
	URL         string              `json:"url"`
	Timestamp   string              `json:"timestamp"` // needs to be ISO8601
	Color       int                 `json:"color"`     // 0xRRGGBB
	Fields      []DiscordEmbedField `json:"fields"`
}

type DiscordEmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

func NewClient() Discord {
	return Discord{
		Token: os.Getenv("DISCORD_API"),
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (d *Discord) SendDiscordMessageFromTFC(ctx context.Context, tfc tfc.TerraformWebhook) {
	log := ctx.Value("log").(*zap.Logger)

	for i, n := range tfc.Notifications {
		discordMsg := DiscordWebhook{
			Content: "",
			Embeds: []DiscordEmbed{
				{
					Title:       "Terraform Status",
					Description: fmt.Sprintf("**%s**", n.Message),
					URL:         tfc.RunURL,
					Color:       getColorForStatus(n.RunStatus),
					Fields:      setFieldsFromRequest(i, tfc),
				},
			},
		}

		jsonBody, err := json.Marshal(discordMsg)
		if err != nil {
			log.Fatal("marshal discord webhook failed",
				zap.Error(err))
		}
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, d.Token+"?wait=true", bytes.NewReader(jsonBody))
		if err != nil {
			log.Fatal("request creation failed",
				zap.Error(err))
		}
		req.Header.Add("content-type", "application/json")

		resp, err := d.HTTPClient.Do(req)
		if err != nil {
			log.Fatal("make discord call",
				zap.Error(err))
		}
		defer resp.Body.Close()
	}
}
