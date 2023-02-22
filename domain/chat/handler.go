package chat

import (
	"fmt"
	"strings"

	"github.com/dragonforce2010/chatgpt-service/constant"
	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
)

const prompt_context_prefix = "The conversaction context:"
const prompt_question_prefix = "Now please only answer my current question:\n"

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

	model := c.Request.URL.Query().Get("model")
	if strings.TrimSpace(model) == "" {
		model = gogpt.GPT3TextDavinci003
	}

	fmt.Println("Received a request: ", chatGptRequest)
	prompt := prompt_context_prefix + chatGptRequest.Context +
		prompt_question_prefix + chatGptRequest.Message + "\n"

	fmt.Printf("current prompt: %v\n", prompt)

	respMessage, err := ch.chatService.GetChatResponse(c, prompt, model)
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
