package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/Jocerdikiawann/server_share_trip/model/request"
)

type TelegramRepositoryImpl struct{}

func NewTelegramRepository() *TelegramRepositoryImpl {
	return &TelegramRepositoryImpl{}
}

func (repo *TelegramRepositoryImpl) SendMessage(context context.Context, chat string) error {
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
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return nil
}
