package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/Jocerdikiawann/server_share_trip/model/request"
	"github.com/sirupsen/logrus"
)

type LoggerHook struct{}

func sendMessageToTelegram(context context.Context, chat string) error {
	baseUrl := os.Getenv("TELEGRAM_URL")
	tokenTelegram := os.Getenv("TELEGRAM_TOKEN")

	groupId, err := strconv.ParseInt(os.Getenv("GROUP_ID"), 10, 64)
	if err != nil {
		return err
	}

	payload := request.MessageTelegram{
		Text:                  chat,
		DisableWebPagePreview: false,
		DisableNotification:   false,
		ChatID:                groupId,
		ParseMode:             "HTML",
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/%s/sendMessage", baseUrl, tokenTelegram)
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	_, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return nil
}

func (l *LoggerHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel, logrus.InfoLevel}
}

func (l *LoggerHook) Fire(entry *logrus.Entry) error {
	eChan := make(chan error)
	go func() {
		chat := fmt.Sprintf("Log %v at %v \n<b>!!The Message!!</b>\n%v", entry.Level, entry.Time, entry.Message)
		if err := sendMessageToTelegram(entry.Context, chat); err != nil {
			eChan <- err
		}
	}()
	close(eChan)
	return <-eChan
}

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.AddHook(&LoggerHook{})

	return logger
}
