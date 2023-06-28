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
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (l *LoggerHook) Fire(entry *logrus.Entry) error {
	l.Repo.SendMessage(entry.Context, fmt.Sprintf("Message : %v, Level : %v, Time: %v", entry.Message, entry.Level, entry.Time))
	return nil
}

func NewLogger(repo design.TelegramRepository) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.AddHook(&LoggerHook{Repo: repo})

	return logger
}
