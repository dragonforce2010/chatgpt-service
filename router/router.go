package router

import (
	"github.com/dragonforce2010/chatgpt-service/domain/chat"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router     *gin.Engine
	chatRouter *chat.ChatRouter
}

func NewRouter(chatRouter *chat.ChatRouter, chatHandler *chat.ChatHandler) *Router {
	return &Router{
		chatRouter: chatRouter,
	}
}

func (r *Router) InitRoutes() {
	r.router = gin.Default()
	r.router.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"POST"},
		AllowCredentials:       true,
	}))
	r.chatRouter.Init(r.router)
}

func (r *Router) Run(address string) {
	r.InitRoutes()
	//// # Headers
	// Allow CORS
	r.router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	})
	r.router.Run(address)
}
