package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"strconv"
)

func printCommandEvents(analyticsChanel <-chan *slacker.CommandEvent) {
	for event := range analyticsChanel {
		fmt.Println("Command Event:")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-")
	os.Setenv("SLACK_APP_TOKEN", "xapp-")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Handler: func(context slacker.BotContext, request slacker.Request, writer slacker.ResponseWriter) {
			year := request.Param("year")
			yob, error := strconv.Atoi(year)
			if error != nil {
				fmt.Println(error)
			}
			age := 2024 - yob
			r := fmt.Sprintf("%d years ago", age)
			writer.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
