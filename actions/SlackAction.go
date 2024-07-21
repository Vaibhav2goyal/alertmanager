package actions

import (
	"alertmanager/models"
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

type SlackAction struct{}

func (a SlackAction) Execute(alert models.Alert) {
	// Send a slack alert
	sendToSlack(alert)
}

func sendToSlack(alert models.Alert) {
	// Get the slack token and channel ID
	token := os.Getenv("SLACK_API_TOKEN")
	channelID := os.Getenv("SLACK_CHANNEL_ID")
	api := slack.New(token)

	// Create Slack message attachment fields
	fields := []slack.AttachmentField{
		{
			Title: "Summary",
			Value: alert.Annotations["summary"],
			Short: false,
		},
		{
			Title: "Description",
			Value: alert.Annotations["description"],
			Short: false,
		},
		{
			Title: "Started At",
			Value: alert.StartsAt,
			Short: true,
		},
		{
			Title: "Status",
			Value: alert.Status,
			Short: true,
		},
	}

	// Add each label as a separate field
	for key, value := range alert.Labels {
		fields = append(fields, slack.AttachmentField{
			Title: key,
			Value: value,
			Short: true,
		})
	}
	//Add preview text
	previewText := fmt.Sprintf(
		"*Alert Summary:* %s\n",
		alert.Annotations["summary"],
	)

	attachment := slack.Attachment{
		Pretext: "🚨 Alert Received",
		Fields:  fields,
		Color:   "#FF0000", // Red color to indicate critical alert
	}
	//Post the message to slack
	_, _, err := api.PostMessage(channelID, slack.MsgOptionText(previewText, false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		fmt.Println("Error sending message to Slack:", err)
	}
}
