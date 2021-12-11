package api

import (
	"encoding/json"
	"net/http"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	b, err := tb.NewBot(tb.Settings{
		Token:       os.Getenv("KIRIRIN_TELEGRAM_TOKEN"),
		Synchronous: true,
	})

	if err != nil {
		panic(err)
	}

	var (
		u    tb.Update
		body []byte
	)

	if _, err := r.Body.Read(body); err != nil {
		panic(err)
	}

	if err = json.Unmarshal(body, &u); err != nil {
		panic(err)
	}

	b.ProcessUpdate(u)
}
