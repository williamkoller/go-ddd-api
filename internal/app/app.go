package app

import (
	"github.com/williamkoller/go-ddd-api/internal/user/handler"
	"github.com/williamkoller/go-ddd-api/internal/user/repository"
	"github.com/williamkoller/go-ddd-api/internal/user/usecase"
	"github.com/williamkoller/go-ddd-api/pkg/config"
	"github.com/williamkoller/go-ddd-api/pkg/database"
)

func Start() {
	config.LoadEnv()

	db, err := database.ConnectPostgres()
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUC)

	r := SetupRouter(userHandler)

	port := config.GetEnv("APP_PORT", "8080")
	r.Run(":" + port)
}
