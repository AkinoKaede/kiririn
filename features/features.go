package features

import (
	"context"

	"github.com/AkinoKaede/kiririn/v2/common/session"
)

type ProcessFunc func(context.Context) interface{}

var features = make(map[interface{}]ProcessFunc)

func RegisterFeature(endpoint interface{}, process ProcessFunc) {
	features[endpoint] = process
}

func Handle(ctx context.Context) {
	b := session.BotFromContext(ctx)

	for endpoint, process := range features {
		b.Handle(endpoint, process(ctx))
	}
}
