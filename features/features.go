package features

import (
	"context"

	"github.com/AkinoKaede/kiririn/v2/common/session"
)

type FeatureFunc func(context.Context) interface{}

var features = make(map[interface{}]FeatureFunc)

func RegisterFeature(endpoint interface{}, feature FeatureFunc) {
	features[endpoint] = feature
}

func Handle(ctx context.Context) {
	b := session.BotFromContext(ctx)

	for endpoint, feature := range features {
		b.Handle(endpoint, feature(ctx))
	}
}
