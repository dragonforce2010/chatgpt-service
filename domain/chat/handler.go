package chat

import (
	"fmt"
	"strings"

	"github.com/dragonforce2010/chatgpt-service/constant"
	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
)

const prompt_context_prefix = "Context:"
const prompt_question_prefix = "Human:"
const prompt_question_postfix = "AI:"

type ChatHandler struct {
	chatService *ChatService
}

func NewChatHandler(chatService *ChatService) *ChatHandler {
	return &ChatHandler{chatService: chatService}
}

func (ch *ChatHandler) HandleChatV1(c *gin.Context) {
	var chatGptRequest ChatGptRequest
	err := c.BindJSON(&chatGptRequest)
	if err != nil {
		c.JSON(constant.HTTPStatusCodeBadRequest, "Invalid request parameter")
		return
	}

	model := gogpt.GPT3TextDavinci003
	if len(chatGptRequest.Model) != 0 {
		model = chatGptRequest.Model
	}

	openAiKey := strings.TrimSpace(chatGptRequest.OpenAiKey)
	if len(openAiKey) == 0 {
		c.JSON(constant.HTTPStatusCodeBadRequest, "OpenAiKey is not provided")
		return
	}

	fmt.Println("Received a request: ", chatGptRequest)
	prompt := prompt_context_prefix + chatGptRequest.Context +
		prompt_question_prefix + chatGptRequest.Message + "\n" +
		prompt_question_postfix + "\n"

	fmt.Printf("current prompt: %v\n", prompt)

	respMessage, err := ch.chatService.Chat(c, gogpt.NewClient(openAiKey), prompt, model, false)
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

func (ch *ChatHandler) HandleChatV2(c *gin.Context) {
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
		prompt_question_prefix + chatGptRequest.Message + "\n" +
		prompt_question_postfix

	fmt.Printf("current prompt: %v\n", prompt)

	respMessage, err := ch.chatService.Chat(c, nil, prompt, model, true)
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
