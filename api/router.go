package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thegrandpackard/PackardQuest/managers"
)

type api struct {
	playerManager managers.PlayerManager
}

func NewServer(playerManager managers.PlayerManager) {
	a := api{
		playerManager: playerManager,
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("api/latest/player/:id", a.getPlayer)
	r.POST("api/latest/player", a.registerPlayer)
	// r.PUT("api/latest/player/:id", a.updatePlayer)

	go r.Run(":8000")
}

func getIntParam(c *gin.Context, valName string) (int, error) {
	valString := c.Param(valName)
	val, err := strconv.Atoi(valString)
	if err != nil {
		return 0, err
	}

	return int(val), nil
}
