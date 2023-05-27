package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	// Print command events received by the bot
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {

	// Set Slack bot token and app token using environment variables
	os.Setenv("SLACK_BOT_TOKEN", "Add you bot Token")
	os.Setenv("SLACK_APP_TOKEN", "Add your app Token")


	// Create a new Slack client with the bot token and app token
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	// Print command events in a separate goroutine
	go printCommandEvents(bot.CommandEvents())

	// Define a command for age calculation
	bot.Command("My old is <year>", &slacker.CommandDefinition{
		Description: "old calculator. Example: My old is 2000",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			old, err := strconv.Atoi(year)
			if err != nil {
				println("error")
				return
			}
			age := 2023 - old
			r := fmt.Sprintf("Your age is %d", age)
			response.Reply(r)
		},
	})

	// Create a context for the bot's event loop
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start listening for events and handle them
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
