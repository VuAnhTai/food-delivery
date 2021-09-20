package main

import (
	"food-delivery/component"
	"food-delivery/modules/restaurant/restauranttransport/ginrestaurant"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	appCtx := component.NewAppContext(db)

	restaurants := r.Group("/restaurants")
	{
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant((appCtx)))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant((appCtx)))

	}
	return r.Run()
}
