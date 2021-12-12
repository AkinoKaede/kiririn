package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/AkinoKaede/kiririn/v2/common"
	"github.com/AkinoKaede/kiririn/v2/common/session"
	"github.com/AkinoKaede/kiririn/v2/features"
	_ "github.com/AkinoKaede/kiririn/v2/main/distro/all"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	ctx = context.Background()
)

func Handler(w http.ResponseWriter, r *http.Request) {
	b := session.BotFromContext(ctx)

	body, err := io.ReadAll(r.Body)
	common.Must(err)

	var u tb.Update
	common.Must(json.Unmarshal(body, &u))

	b.ProcessUpdate(u)
}

func init() {
	b, err := tb.NewBot(tb.Settings{
		Token:       os.Getenv("KIRIRIN_TELEGRAM_TOKEN"),
		Synchronous: true,
	})
	common.Must(err)

	ctx = session.ContextWithBot(ctx, b)
	features.Handle(ctx)
}
