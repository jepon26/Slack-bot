package main

import (
	"context"
	"fmt"
	"log"
	"logo"
	"os"

	"github.com/shomali11/slacker"
	"golang.org/x/tools/go/analysis"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Comand Events")
		fmt.Fprintln(event.Timestamp)
		fmt.Fprintln(event.Comand)
		fmt.Println(event.Paramaters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "")
	os.Setenv("SLACK_APP_TOKEN", "")

	bot := slacker.NewClient(os.Getenv("SLAcK_BOT_TOCKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())


	bot.Comand("My age is <year>", &slacker.CommandDefinition){
		Description := "age calculator",
		Example := "My age is 1900",
		Handler: func (botCtx slacker.botContext, request slacker.Request, response slacker.Response)  {
			
		}
	}


	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.listen(ctx)
	if err != nil{
		log.Fatal(err)
	}
}
