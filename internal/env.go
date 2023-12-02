package internal

import "time"

type Env struct {
	ApiURL   string
	ApiToken string

	BotToken string
	ChatId   int64

	StationFrom int
	StationTo   int
	DateFrom    time.Time
	DateTo      time.Time
}

func NewEnv(
	apiURL string,
	apiToken string,
	botToken string,
	chatId int64,
	stationFrom int,
	stationTo int,
	dateFrom time.Time,
	dateTo time.Time,
) *Env {
	return &Env{
		ApiURL:      apiURL,
		ApiToken:    apiToken,
		BotToken:    botToken,
		ChatId:      chatId,
		StationFrom: stationFrom,
		StationTo:   stationTo,
		DateFrom:    dateFrom,
		DateTo:      dateTo,
	}
}
