package main

import(
	"fmt"
	"github.com/slack-go/slack"
	"github.com/Shimodaira0910/summary/env"
)

func main(){	
	env := env.Env{}
	api := slack.New(env.LoadEnv("SLACK_BOT_TOKEN"))
	channelId := env.LoadEnv("SLACK_CHANNEL_ID")

	channelId, timestamp, err := api.PostMessage(
		channelId,
		slack.MsgOptionText("hello!!", false),
	)
	if err != nil{
		fmt.Println("接続失敗")
		fmt.Println(err)
		return
	}

	fmt.Printf("Message successfully sent to channel %s at %s", channelId, timestamp)
}	