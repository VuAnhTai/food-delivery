package rstlikebiz

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/modules/restaurantlike/model"
	"food-delivery/pubsub"
)

type UserUnlikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

type userUnlikeRestaurantBiz struct {
	store  UserUnlikeRestaurantStore
	pubsub pubsub.Pubsub
}

func NewUserUnlikeRestaurantBiz(store UserUnlikeRestaurantStore, pubsub pubsub.Pubsub) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{store: store, pubsub: pubsub}
}

func (biz *userUnlikeRestaurantBiz) UnlikeRestaurant(
	ctx context.Context,
	userId,
	restaurantId int,
) error {
	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCannotUnlikeRestaurant(err)
	}

	// side effect
	biz.pubsub.Publish(ctx, common.TopicUserDislikeRestaurant, pubsub.NewMessage(&restaurantlikemodel.Like{
		RestaurantId: restaurantId,
		UserId:       userId,
	}))

	return nil
}
