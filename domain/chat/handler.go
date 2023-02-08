package chat

import (
	"fmt"
	"github.com/dragonforce2010/chatgpt-service/constant"
	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	chatService *ChatService
}

func NewChatHandler(chatService *ChatService) *ChatHandler {
	return &ChatHandler{chatService: chatService}
}
func (ch *ChatHandler) HandleChat(c *gin.Context) {
	var chatGptRequest ChatGptRequest
	err := c.BindJSON(&chatGptRequest)
	if err != nil {
		c.JSON(constant.HTTPStatusCodeBadRequest, "Invalid request parameter")
		return
	}

	fmt.Println("Received a request: ", chatGptRequest)
	respMessage, err := ch.chatService.GetChatResponse(c, chatGptRequest.Message)
	if err != nil {
		c.JSON(constant.HTTPStatusCodeInternalError, ChatGptResponse{
			Content: "",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(constant.HTTPStatusCodeSuccess, ChatGptResponse{
		Content: respMessage,
	})
}
