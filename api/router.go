package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/thegrandpackard/PackardQuest/managers"
)

type api struct {
	playerManager         managers.PlayerManager
	triviaQuestionManager managers.TriviaQuestionManager

	upgrader websocket.Upgrader
	clients  map[int]*websocket.Conn
}

func NewServer(playerManager managers.PlayerManager, triviaQuestionmanager managers.TriviaQuestionManager) {
	a := &api{
		playerManager:         playerManager,
		triviaQuestionManager: triviaQuestionmanager,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
		clients: map[int]*websocket.Conn{},
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

	r.GET("ws/player/:id", a.websocketUpgraderPlayer)

	apiLatest := r.Group("api/latest")

	// player
	apiLatest.GET("player/:id", a.getPlayer)
	apiLatest.POST("player", a.registerPlayer)
	apiLatest.PUT("player/:id", a.updatePlayer)

	// trivia
	apiLatest.GET("trivia/player/:id", a.getPlayerTriviaQuestion)
	apiLatest.POST("trivia/player/:id", a.answerTriviaQuestion)

	// scoreboard
	apiLatest.GET("scoreboard", a.getScoreboard)

	playerManager.SetSubscriber(a)
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
