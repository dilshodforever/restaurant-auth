package api

import (
	"github.com/dilshodforever/restaurant-auth/api/handler"
	"github.com/dilshodforever/restaurant-auth/api/middleware"
	_ "github.com/dilshodforever/restaurant-auth/docs"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title 		Voting service
// @BasePath 	/
// @securityDefinitions.apikey BearerAuth
// @in 			header
// @name 		Authorization
func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}
	r.Use(cors.New(corsConfig))

	r.Use(middleware.MiddleWare())

	u := r.Group("/user")
	u.POST("/create", h.RegisterUser)
	u.PUT("/update/:id", h.UpdateUser)
	u.DELETE("/delete/:id", h.DeleteUser)
	u.GET("/getall", h.GetAllUser)
	u.GET("/getbyid/:id", h.GetbyIdUser)
	u.POST("/login", h.LoginUser)
	
	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger. WrapHandler(files.Handler, url))
	return r
}
