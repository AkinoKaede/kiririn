package about

import (
	"context"
	"fmt"

	"github.com/AkinoKaede/kiririn/v2/common/session"
	"github.com/AkinoKaede/kiririn/v2/features"
	"github.com/AkinoKaede/kiririn/v2/kiririn"

	tb "gopkg.in/tucnak/telebot.v2"
)

func Process(ctx context.Context) interface{} {
	b := session.BotFromContext(ctx)

	return func(m *tb.Message) {
		b.Reply(m, fmt.Sprintf("Kiririn v%s - %s", kiririn.Version(), kiririn.Usage()))
	}
}

func init() {
	features.RegisterFeature("/about", Process)
}
