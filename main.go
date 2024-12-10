package main

import (
	"flag"
	"fmt"

	"github.com/ToffaKrtek/go-tg-cli/environment"
	"github.com/ToffaKrtek/go-tg-cli/send"
)

var envPath = ""

func main() {
	environment.Parse(envPath)
	flagMsg := flag.String("msg", "", "Message from cli")
	flagFile := flag.String("file", "", "File path")
	flagImage := flag.String("image", "", "Image path")
	flagType := flag.String("type", "", "Type: [ success, warning, error ]")
	flagParseMode := flag.String("mode", "", "Telegram parse mode")
	flagChatId := flag.String("chat_id", environment.Get("TG_CHAT_ID"), "Telegram chat ID")
	flagTopicId := flag.String("topic_id", environment.Get("TG_TOPIC_ID"), "Telegram topic ID")
	flagToken := flag.String("token", environment.Get("TG_TOKEN"), "Telegram bot token")

	flag.Parse()

	message := send.NewMessage(
		*flagToken,
		send.Text(*flagMsg),
		send.Image(*flagImage),
		send.File(*flagFile),
		send.ParseMode(*flagParseMode),
		send.Type(*flagType),
		send.ChatId(*flagChatId),
		send.Topic(*flagTopicId),
	)
	if err := message.Send(); err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
	fmt.Println("Message sent successfully")
}
