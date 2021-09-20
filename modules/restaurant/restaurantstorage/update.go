package restaurantstorage

import (
	"context"
	"food-delivery/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) UpdateDataByCondition(
	ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
