package chat

import (
	"fmt"

	"github.com/dragonforce2010/chatgpt-service/constant"
	"github.com/gin-gonic/gin"
)

const prompt_prefix = "ask:\n"
const prompt_postfix = "answer:\n"

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
	prompt := prompt_prefix +
		chatGptRequest.Context + "\n" +
		chatGptRequest.Message + "\n" +
		prompt_postfix
	fmt.Printf("prompt: %v\n", prompt)

	respMessage, err := ch.chatService.GetChatResponse(c, prompt)
	if err != nil {
		c.JSON(constant.HTTPStatusCodeInternalError, ChatGptResponse{
			Content: "",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(constant.HTTPStatusCodeSuccess, ChatGptResponse{
		Content: respMessage,
		Context: chatGptRequest.Context + chatGptRequest.Message + "\n",
	})
}
