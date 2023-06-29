package utils

import (
	"fmt"

	"github.com/Jocerdikiawann/server_share_trip/repository/design"
	"github.com/sirupsen/logrus"
)

type LoggerHook struct {
	Repo design.TelegramRepository
}

func (l *LoggerHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel, logrus.InfoLevel}
}

func (l *LoggerHook) Fire(entry *logrus.Entry) error {
	eChan := make(chan error)
	go func() {
		chat := fmt.Sprintf("Log %v at %v \n<b>!!The Message!!</b>\n%v", entry.Level, entry.Time, entry.Message)
		if err := l.Repo.SendMessage(entry.Context, chat); err != nil {
			eChan <- err
		}
	}()
	close(eChan)
	return <-eChan
}

func NewLogger(repo design.TelegramRepository) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.AddHook(&LoggerHook{Repo: repo})

	return logger
}
