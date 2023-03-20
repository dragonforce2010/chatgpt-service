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
	// 使用自己的open ai key
	router.POST("/api/chat/v1", cr.handler.HandleChatV1)
	// 使用public open api key pool
	router.POST("/api/chat/v2", cr.handler.HandleChatV2)
	// router.POST("/api/chat/chatgpt-demo-for-apaas", cr.handler.HandleChatV2)
}
