package handler

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/yanzay/tbot/v2"
)

func Main() {
	ctx := context.Background()
	bot := tbot.New(os.Getenv("TELEGRAM_BOT_TOKEN"),
		tbot.WithWebhook("https://golang-chat-gpt-telegram-bot-vercel.vercel.app", ":8080"))
	c := bot.Client()

	bot.HandleMessage(".*human:*", func(m *tbot.Message) {
		///////
		c0 := gogpt.NewClient(os.Getenv("OPENAI_TOKEN"))

		maxtokens, err0 := strconv.Atoi(os.Getenv("OPENAI_MAXTOKENS"))

		if err0 != nil {
			fmt.Println("Error during conversion")
			//return "MaxTokens Conversion Error happened!"
		}

		req := gogpt.CompletionRequest{
			Model:       "text-davinci-003",
			MaxTokens:   maxtokens,
			Prompt:      m.Text, //the text you typed in with human:.....
			Temperature: 0,
		}
		resp, err := c0.CreateCompletion(ctx, req)
		if err != nil {
			//return "You got an error!"
		} else {
			fmt.Println(resp.Choices[0].Text)

			//return resp.Choices[0].Text
		}

		////////////

		c.SendMessage(m.Chat.ID, "AI:"+resp.Choices[0].Text)
	})
	log.Fatal(bot.Start())
}
