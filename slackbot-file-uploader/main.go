package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	tokenErr := os.Setenv("SLACK_BOT_TOKEN", "")
	if tokenErr != nil {
		return
	}
	channelErr := os.Setenv("CHANNEL_ID", "")
	if channelErr != nil {
		return
	}

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"ZIPL.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			File:     fileArr[i],
			Channels: channelArr,
		}

		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
