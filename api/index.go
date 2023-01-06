package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/yanzay/tbot/v2"
)

// HandlerFunc is the signature type for the main function that will handle HTTP requests.
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of the GPT-3 client.
	c := gogpt.NewClient(os.Getenv("OPENAI_TOKEN"))
	ctx := context.Background()

	bot := tbot.New(os.Getenv("TELEGRAM_BOT_TOKEN"),
		tbot.WithWebhook("https://golang-chat-gpt-telegram-bot-vercel.vercel.app/", ":8080"))
	c1 := bot.Client() //Please add your Render URL between "". 請在引號中加入你的Render網址

	/////////////////

	bot.HandleMessage(".*human:*", func(m *tbot.Message) { //The AI must be triggered by the keyword "human:"

		req := gogpt.CompletionRequest{
			Model:     "text-davinci-003",
			MaxTokens: 200,
			Prompt:    m.Text, //the text you typed in
		}
		resp, err := c.CreateCompletion(ctx, req)
		if err != nil {
			return
		}
		fmt.Println(resp.Choices[0].Text) // the answer you got

		c1.SendMessage(m.Chat.ID, "AI:"+resp.Choices[0].Text) //m.Text represents the text you typed in 代表你打的文字
	})
	log.Fatal(bot.Start())
}
