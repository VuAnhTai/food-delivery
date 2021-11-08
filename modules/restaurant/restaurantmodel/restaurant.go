package restaurantmodel

import (
	"errors"
	"food-delivery/common"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`   // inline return flat if no => create new object
	Name            string             `json:"name" gorm:"column:name;"`
	UserId          int                `json:"-" gorm:"column:owner_id;"`
	Addr            string             `json:"address" gorm:"column:addr;"`
	Logo            *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images     `json:"cover" gorm:"column:cover;"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false;"`
	LikedCount      int                `json:"liked_count" gorm:"column:liked_count;"` // computed field
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"` // inline return flat if no => create new object
	Name            string           `json:"name" gorm:"column:name;"`
	UserId          int              `json:"-" gorm:"column:owner_id;"`
	Addr            string           `json:"address" gorm:"column:addr;"`
	Logo            *common.Image    `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images   `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("restaurant name can't be blank")
	}

	return nil
}

type RestaurantUpdate struct {
	Name  string         `json:"name" gorm:"column:name;"`
	Addr  string         `json:"address" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func (data *Restaurant) Mask(isAdminOrOwner bool) {
	data.GenUid(common.DbTypeRestaurant)

	if u := data.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}
