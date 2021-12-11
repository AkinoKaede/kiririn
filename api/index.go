package api

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/AkinoKaede/kiririn/v2/common"
	"github.com/AkinoKaede/kiririn/v2/common/session"
	"github.com/AkinoKaede/kiririn/v2/features"

	tb "gopkg.in/tucnak/telebot.v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	b, err := tb.NewBot(tb.Settings{
		Token:       os.Getenv("KIRIRIN_TELEGRAM_TOKEN"),
		Synchronous: true,
	})
	common.Must(err)

	ctx := session.ContextWithBot(context.Background(), b)

	features.Handle(ctx)

	var (
		u    tb.Update
		body []byte
	)

	common.Must2(r.Body.Read(body))
	common.Must(json.Unmarshal(body, &u))

	b.ProcessUpdate(u)
}
