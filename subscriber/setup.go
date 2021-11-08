package subscriber

import (
	"context"
	"food-delivery/component"
)

func Setup(ctx component.AppContext) {
	IncreaseLikeCountAfterUserLikeRestaurant(ctx, context.Background())
}
