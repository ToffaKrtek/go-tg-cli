package send

import (
	"encoding/json"
	"strings"
)

type Message struct {
	BaseUrl   string
	Text      string
	Image     string
	File      string
	ParseMode string
	Type      string
	ChatId    string
	Topic     string
}

func NewMessage(token string, opts ...MessageFunc) *Message {
	message := &Message{
		BaseUrl: getTgAPi(token),
	}
	for _, opt := range opts {
		opt(message)
	}
	return message
}

func (m Message) getMsgData(text string, mode string) []byte {
	res := map[string]string{
		"chat_id":    m.ChatId,
		"text":       text,
		"parse_mode": mode,
	}
	if len(m.Topic) > 0 {
		res["topic_id"] = m.Topic
	}
	data, err := json.Marshal(res)
	if err != nil {
		return []byte{}
	}
	return data
}

type MessageFunc func(*Message)

func ChatId(chatId string) MessageFunc {
	return func(m *Message) {
		m.ChatId = chatId
	}
}

func Topic(topic string) MessageFunc {
	return func(m *Message) {
		if len(topic) > 0 {
			m.Topic = topic
		}
	}
}
func Text(text string) MessageFunc {
	return func(m *Message) {
		m.Text = text
	}
}

func Image(path string) MessageFunc {
	return func(m *Message) {
		m.Image = path
	}
}

func File(path string) MessageFunc {
	return func(m *Message) {
		m.File = path
	}
}

func ParseMode(mode string) MessageFunc {
	return func(m *Message) {
		var isValid bool
		modeLower := strings.ToLower(mode)
		for _, v := range parseModes {
			if v == modeLower {
				isValid = true
			}
		}
		if isValid {
			m.ParseMode = modeLower
		}
	}
}

func Type(t string) MessageFunc {
	return func(m *Message) {
		var isValid bool
		typeLower := strings.ToLower(t)
		for _, v := range statusTypes {
			if v == typeLower {
				isValid = true
			}
		}
		if isValid {
			m.Type = typeLower
		}
	}
}
