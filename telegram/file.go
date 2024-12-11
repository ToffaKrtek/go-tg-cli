package telegram

import (
	"fmt"
	"os"

	"github.com/ToffaKrtek/go-tg-cli/s3"
)

const maxfilesize = 50 * 1024 * 1024

func (m Message) makeSendFile() error {
	isLarge, err := isLargeFile(m.File)
	if err != nil {
		return err
	}
	if isLarge {
		f := s3.NewFile(m.File)

		link, err := f.Upload()
		if err != nil {
			return err
		}
		msg := fmt.Sprintf("<a href='%s'>[Ссылка для скачивания]</a>", link)
		return m.sendTextRequest(m.BaseUrl+"/sendMessage", m.getMsgData(msg, "HTML"))
	}
	return m.sendFile()
}

func isLargeFile(filepath string) (bool, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}
	fileSize := fileInfo.Size()
	if fileSize > maxfilesize {
		return true, nil
	}
	return false, nil
}
