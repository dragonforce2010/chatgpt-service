package chat

import (
	"github.com/gin-gonic/gin"
)

type ChatRouter struct {
	handler *ChatHandler
}

func NewChatRouter(chatHandler *ChatHandler) *ChatRouter {
	return &ChatRouter{
		handler: chatHandler,
	}
}

func (cr *ChatRouter) Init(router *gin.Engine) {
	router.POST("/api/chat", cr.handler.HandleChat)
}
