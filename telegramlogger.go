package Logi

import (
	"strconv"
	"time"

	tele "gopkg.in/telebot.v3"
)

type TelegramLogger struct {
	token     string
	channelID string
	bot       *tele.Bot
	chat      *tele.Chat
}

func NewTelegramLogger(token string, channelID string) Logger {
	logger := TelegramLogger{}

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		panic(err)
	}

	id, err := strconv.ParseInt(channelID, 10, 64)
	if err != nil {
		panic(err)
	}

	chat, err := b.ChatByID(id)
	if err != nil {
		panic(err)

	}

	logger.channelID = channelID
	logger.token = token
	logger.bot = b
	logger.chat = chat

	return &logger
}

func (logger TelegramLogger) sendMessage(message string) {
	_, err := logger.bot.Send(logger.chat, message)
	if err != nil {
		panic(err)
	}
}

func (logger TelegramLogger) Log(message string) {
	logger.sendMessage(message)
}

func (logger TelegramLogger) Warning(message string) {
	logger.sendMessage(message)
}

func (logger TelegramLogger) Error(message string) {
	logger.sendMessage(message)
}
