package features

import (
	"context"

	"github.com/AkinoKaede/kiririn/v2/common/session"
	"github.com/AkinoKaede/kiririn/v2/features/about"
)

type ProcessFunc func(context.Context) interface{}

var features = map[interface{}]ProcessFunc{
	"/about": about.Process,
}

func Handle(ctx context.Context) {
	b := session.BotFromContext(ctx)

	for endpoint, process := range features {
		b.Handle(endpoint, process(ctx))
	}
}
