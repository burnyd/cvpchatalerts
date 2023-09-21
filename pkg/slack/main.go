package slack

import (
	"log"
	"os"
	"time"

	"cvpchatalert/pkg/jsonlog"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func SendtoSlack(alertTitle, alertResponse string) {

	godotenv.Load()

	token := os.Getenv("SLACKTOK")
	if token == "" {
		log.Fatalln("Missing API KEY")
	}

	channelID := os.Getenv("SLACKCHAN")
	if token == "" {
		log.Fatalln("Missing API KEY")
	}

	client := slack.New(token, slack.OptionDebug(true))
	// Create the Slack attachment that we will send to the channel
	attachment := slack.Attachment{
		Pretext: "Alert from CVP",
		Title:   alertTitle,
		Text:    alertResponse,
		// Color Styles the Text, making it possible to have like Warnings etc.
		Color:     "#36a64f",
		TitleLink: "https://10.255.62.214/cv/events",
		// Fields are Optional extra data!
		Fields: []slack.AttachmentField{
			{
				Title: "Date",
				Value: time.Now().String(),
			},
		},
	}
	// PostMessage will send the message away.
	// First parameter is just the channelID, makes no sense to accept it
	_, timestamp, err := client.PostMessage(
		channelID,
		// uncomment the item below to add a extra Header to the message, try it out :)
		//slack.MsgOptionText("New message from bot", false),
		slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		panic(err)
	}

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	logger.PrintInfo("Message sent at "+timestamp, nil)
}
