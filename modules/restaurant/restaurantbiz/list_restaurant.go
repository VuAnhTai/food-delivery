package restaurantbiz

import (
	"context"
	"food-delivery/common"
	"food-delivery/modules/restaurant/restaurantmodel"
	"log"
)

type ListRestaurantStore interface {
	ListDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type LikeStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}
type listRestaurantBiz struct {
	store     ListRestaurantStore
	likeStore LikeStore
}

func NewListRestaurantBiz(store ListRestaurantStore, likeStore LikeStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store, likeStore: likeStore}
}

func (biz *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	data, err := biz.store.ListDataByCondition(
		ctx, nil, filter, paging, "User",
	)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	ids := make([]int, len(data))

	for i := range data {
		ids[i] = data[i].Id
	}

	mapResLike, err := biz.likeStore.GetRestaurantLikes(ctx, ids)

	if err != nil {
		log.Println(err)
	}

	if v := mapResLike; v != nil {
		for i, item := range data {
			data[i].LikedCount = mapResLike[item.Id]
		}
	}
	return data, err
}
