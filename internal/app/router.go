package app

import (
	"github.com/gin-gonic/gin"
	"github.com/williamkoller/go-ddd-api/internal/auth"
	"github.com/williamkoller/go-ddd-api/internal/user/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/williamkoller/go-ddd-api/docs"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/users/register", userHandler.Register)
	r.POST("/users/login", userHandler.Login)

	authMiddleware := auth.JWTAuthMiddleware()
	userGroup := r.Group("/users")
	userGroup.Use(authMiddleware)
	{
		userGroup.GET("/", userHandler.List)
		userGroup.PUT("/:id", userHandler.Update)
		userGroup.DELETE("/:id", userHandler.Delete)
	}

	return r
}
