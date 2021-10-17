package discord

import "tfnotifications/pkg/tfc"

const (
	nullString = "(null)"
	red        = 0xff0000
	green      = 0x00ff00
	yellow     = 0xedb021
	blue       = 0x3b6bed

	planned_and_finishedString = "planned_and_finished"
	appliedString              = "applied"
	erroredString              = "errored"
	plannedString              = "planned"
)

func getColorForStatus(status string) int {

	switch status {
	case planned_and_finishedString:
		return green
	case appliedString:
		return green
	case erroredString:
		return red
	case plannedString:
		return blue
	default:
		return yellow
	}
}

func setFieldsFromRequest(iterator int, tfc tfc.TerraformWebhook) []DiscordEmbedField {
	var fields []DiscordEmbedField
	if tfc.Notifications[iterator].RunStatus == "" {
		tfc.Notifications[iterator].RunStatus = nullString
	}

	if tfc.RunMessage == "" {
		tfc.RunMessage = nullString
	}

	if tfc.RunCreatedBy == "" {
		tfc.RunCreatedBy = nullString
	}

	if tfc.Notifications[iterator].RunUpdatedBy == "" {
		tfc.Notifications[iterator].RunUpdatedBy = nullString
	}

	fields = []DiscordEmbedField{
		{
			Name:   "Run Status",
			Value:  tfc.Notifications[iterator].RunStatus,
			Inline: false,
		},
		{
			Name:   "Run Message",
			Value:  tfc.RunMessage,
			Inline: false,
		},
		{
			Name:   "Run Created By",
			Value:  tfc.RunCreatedBy,
			Inline: true,
		},
		{
			Name:   "Run Updated By",
			Value:  tfc.Notifications[iterator].RunUpdatedBy,
			Inline: true,
		},
		{
			Name:   "Trigger",
			Value:  tfc.Notifications[iterator].Trigger,
			Inline: true,
		},
	}
	return fields
}
