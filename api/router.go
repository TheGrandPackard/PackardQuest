package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/thegrandpackard/PackardQuest/interfaces"
)

type api struct {
	playerManager         interfaces.PlayerManager
	triviaQuestionManager interfaces.TriviaQuestionManager
	scoreboardManager     interfaces.ScoreboardManager

	upgrader                    websocket.Upgrader
	playerWebscocketConnections map[int]*websocket.Conn
	pensieveWebsocketConnection *websocket.Conn
}

func NewServer(
	playerManager interfaces.PlayerManager,
	triviaQuestionmanager interfaces.TriviaQuestionManager,
	scoreboardManager interfaces.ScoreboardManager,
) {
	a := &api{
		playerManager:         playerManager,
		triviaQuestionManager: triviaQuestionmanager,
		scoreboardManager:     scoreboardManager,

		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
		playerWebscocketConnections: map[int]*websocket.Conn{},
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

	// websockets
	r.GET("ws/player/:id", a.websocketUpgraderPlayer)
	r.GET("ws/pensieve", a.websocketUpgraderPensieve)

	apiLatest := r.Group("api/latest")

	// player
	apiLatest.GET("player/:id", a.getPlayerByID)
	apiLatest.GET("player/wand/:wandID", a.getPlayerByWandID)
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
