package ginresturantlike

import (
	"food-delivery/common"
	"food-delivery/component"
	"food-delivery/modules/restaurant/restaurantstorage"
	rstlikebiz "food-delivery/modules/restaurantlike/biz"
	restaurantlikestorage "food-delivery/modules/restaurantlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DELETE /v1/restaurants/:id/unlike

func UserUnlikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		//data := restaurantlikemodel.Like{
		//	RestaurantId: int(uid.GetLocalID()),
		//	UserId:       requester.GetUserId(),
		//}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		decStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rstlikebiz.NewUserUnlikeRestaurantBiz(store, decStore)

		if err := biz.UnlikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
