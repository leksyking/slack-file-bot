package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	{
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading env variables")
		}
	}

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
}
