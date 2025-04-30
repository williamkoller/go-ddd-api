package main

import (
	"github.com/williamkoller/go-ddd-api/internal/app"
	"log"
	"runtime"
	"time"
)

// @title Go DDD API
// @version 0.0.1
// @description API com DDD, Auth e CRUD de usuário
// @host localhost:3003
// @BasePath /
// @schemes http, https
func main() {
	go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			log.Printf("[Monitor] Memória usada: %.2f MB", float64(m.Alloc)/1024.0/1024.0)
			time.Sleep(1 * time.Second)
		}
	}()
	app.Start()
}
