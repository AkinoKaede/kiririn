package session

import (
	"context"

	tb "gopkg.in/tucnak/telebot.v2"
)

type sessionKey int

const (
	botSessionKey sessionKey = iota
)

func ContextWithBot(ctx context.Context, bot *tb.Bot) context.Context {
	return context.WithValue(ctx, botSessionKey, bot)
}

func BotFromContext(ctx context.Context) *tb.Bot {
	return ctx.Value(botSessionKey).(*tb.Bot)
}
