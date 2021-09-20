package restaurantbiz

import (
	"context"
	"food-delivery/common"
	"food-delivery/modules/restaurant/restaurantmodel"
)

type UpdateRestaurantStore interface {
	UpdateDataByCondition(
		ctx context.Context,
		id int,
		data *restaurantmodel.RestaurantUpdate,
	) error
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(
	ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	if err := biz.store.UpdateDataByCondition(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err)
	}

	return err
}
