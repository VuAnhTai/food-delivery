package ginuser

import (
	"food-delivery/common"
	"food-delivery/component"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
