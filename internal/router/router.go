package router

import (
	"net/http"

	userservice "github.com/artyomkorchagin/tz-go-gin/internal/services/user"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	_ "github.com/artyomkorchagin/tz-go-gin/docs"
)

type Handler struct {
	userService *userservice.Service
	logger      *zap.Logger
}

func NewHandler(userService *userservice.Service, logger *zap.Logger) *Handler {
	return &Handler{
		userService: userService,
		logger:      logger,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	main := router.Group("/")
	{
		main.POST("/users", h.wrap(h.createUser))
		main.GET("/users/:id", h.wrap(h.getUser))

		main.GET("/status", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		main.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	h.logger.Info("Routes initialized")
	return router
}
