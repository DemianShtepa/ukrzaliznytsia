package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
	"ukrzaliznytsia/internal"
)

func parseEnv() *internal.Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	stationFrom, _ := strconv.Atoi(os.Getenv("STATION_FROM"))
	stationTo, _ := strconv.Atoi(os.Getenv("STATION_TO"))

	dateFrom, _ := time.Parse("2006-01-02", os.Getenv("DATE_FROM"))
	dateTo, _ := time.Parse("2006-01-02", os.Getenv("DATE_TO"))

	chatId, _ := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)

	return internal.NewEnv(
		os.Getenv("API_URL"),
		os.Getenv("API_TOKEN"),
		os.Getenv("BOT_TOKEN"),
		chatId,
		stationFrom,
		stationTo,
		dateFrom,
		dateTo,
	)
}

func main() {
	env := parseEnv()

	events := make(chan time.Time)

	api := internal.NewApi(*env)
	api.Fetch(events, time.Minute*30)

	bot := internal.NewBot(*env)

	for event := range events {
		bot.Notify(event)
	}
}
