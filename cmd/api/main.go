package main

import (
	"fmt"
	"todo_api/internal/config"
	"todo_api/internal/database"
	"todo_api/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	var cfg *config.Config
	var err error

	cfg, err = config.Load()

	if err != nil {
		fmt.Println("Unable to load our configuration")
	}

	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)

	if err != nil {
		fmt.Println("Failed to connect to database")
	}

	defer pool.Close()

	var router *gin.Engine = gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context) {
		//map[string]interface{}
		// map[string]any{}
		c.JSON(200, gin.H{
			"message":  "Todo API is running!",
			"status":   "success",
			"database": "database connected",
		})
	})

	router.POST("/todos",handlers.CreateTodoHandler(pool))
	// start the server
	router.Run(":" + cfg.Port)
}
