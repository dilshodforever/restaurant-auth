package api

import (
	"github.com/dilshodforever/restaurant-auth/api/handler"
	//_ "github.com/dilshodforever/restaurant-auth/api/docs"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Voting service
// @version 1.0
// @description Voting service
// @
// @host localhost:8080
// @BasePath /
func NewGin(h *handler.Handler)*gin.Engine{
	r:=gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
	u:=r.Group("/user")
	u.POST("/login", h.RegisterUser)
	u.PUT("/update/:id", h.UpdateUser)
	u.DELETE("/delete/:id", h.DeleteUser)
	u.GET("/getall", h.GetAllUser)
	u.GET("/getbyid/:id", h.GetbyIdUser)
	
	
	return r
}