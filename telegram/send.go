package telegram

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

var parseModes = []string{
	"HTML",
	"Markdown",
}

var statusTypes = []string{
	"success",
	"warning",
	"error",
}

func (m Message) Send() error {
	if len(m.File) > 0 {
		err := m.sendFile()
		if err != nil {
			return err
		}
	}
	if len(m.Image) > 0 {
		err := m.sendImage()
		if err != nil {
			return err
		}
	}
	if len(m.Text) > 0 {
		err := m.sendText()
		if err != nil {
			return err
		}
	}
	return nil
}

func (m Message) sendImage() error {
	return m.sendFileRequest(m.BaseUrl+"/sendPhoto", m.Image, "photo")
}

func (m Message) sendFile() error {
	return m.sendFileRequest(m.BaseUrl+"/sendDocument", m.File, "document")
}

func (m Message) sendText() error {
	return m.sendTextRequest(m.BaseUrl+"/sendMessage", m.getMsgData(m.Text, m.ParseMode))
}

func (m Message) sendTextRequest(url string, data []byte) error {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (m Message) sendFileRequest(url string, filePath string, partName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(partName, filePath)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	writer.WriteField("chat_id", m.ChatId)

	if len(m.Topic) > 0 {
		writer.WriteField("messaage_thread_id", m.Topic)
	}
	writer.Close()

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	return nil
}
