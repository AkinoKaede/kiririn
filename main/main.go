package main

import (
	"context"
	"os"
	"time"

	"github.com/AkinoKaede/kiririn/v2/common"
	"github.com/AkinoKaede/kiririn/v2/common/session"
	"github.com/AkinoKaede/kiririn/v2/features"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("KIRIRIN_TELEGRAM_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	common.Must(err)

	ctx := session.ContextWithBot(context.Background(), b)

	features.Handle(ctx)

	b.Start()
}
