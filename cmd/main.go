package main

import "github.com/williamkoller/go-ddd-api/internal/app"

// @title Go DDD API
// @version 0.0.1
// @description API com DDD, Auth e CRUD de usuário
// @host localhost:3003
// @BasePath /
// @schemes http, https
func main() {
	app.Start()
}
