package design

import (
	"context"
)

type TelegramRepository interface {
	SendMessage(context context.Context, chat string) error
}
