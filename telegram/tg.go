package telegram

type TG struct {
	Token string
}

var telegramAPI = "https://api.telegram.org/bot"

func getTgAPi(token string) string {
	return telegramAPI + token
}
