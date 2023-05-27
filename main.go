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
<<<<<<< HEAD
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5319951951606-5326648838450-nnkFV5RK1m5g7BNWhMVq15Ay")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A059HJAUCAZ-5335044912292-d7cf89f0401a1c64c7843093f82ebe481bbc5af4958ebeea00fec8fbd1120ba3")
=======
	os.Setenv("SLACK_BOT_TOKEN", "")
	os.Setenv("SLACK_APP_TOKEN", "")
>>>>>>> 5b1391227489dd9618224dbabb50c72a544fc81f

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("My old is <year>", &slacker.CommandDefinition{
		Description: "age calculator. Example: My old is 2000",
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




	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}


