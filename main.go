package main

import (
	"context"
	"fmt"
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
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5319951951606-5326648838450-DaNeshDtnSGsDHr3tBh8MeXI")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A059HJAUCAZ-5323628079893-b238bc91d1a9659d6a9d2e8085b8f581d3995bb7f8db60fdcc90da30da851d89")

	bot := slacker.NewClient(os.Getenv("SLAcK_BOT_TOCKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
}
