package router

import (
	"github.com/babulal107/go-cloud-native-app/internal/config"
	"github.com/babulal107/go-cloud-native-app/internal/handler"
	"github.com/babulal107/go-cloud-native-app/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(appContainer config.AppContainer) *gin.Engine {

	// create default gin router
	r := gin.Default()

	// add default gin middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Ok",
		})
	})

	userSvc := service.NewUserSvc(appContainer)
	userHandler := handler.NewUserHandler(userSvc)

	// Create a group with the prefix "/api/v1"
	api := r.Group("/api/v1")
	{
		api.GET("/users", userHandler.GetUsers)
		api.POST("/users", userHandler.PostUser)
		api.GET("/user/:id", userHandler.GetUser)
	}

	api.POST("/register", userHandler.Register())

	return r

}
