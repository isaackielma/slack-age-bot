package main

import(
	"context"
	"fmt"
	"log"
	"os"
	
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan*slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3637593783267-3872697546516-P0uGLiXjoFq8cX8ZJDzVDjtj")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03R58CT1K9-3869788535091-31234e694cd5f23e1575c29650ab78a4f2a2877b4583f18ebea7964d0df3a1ac")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}
}