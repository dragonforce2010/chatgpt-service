package chat

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dragonforce2010/chatgpt-service/constant"
	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
)

const ROLE_AI = "AI"
const ROLE_USER = "USER"
const MAX_CONTEXT_SIZE = 20

type ChatHandler struct {
	chatService *ChatService
}

func NewChatHandler(chatService *ChatService) *ChatHandler {
	return &ChatHandler{chatService: chatService}
}

func (ch *ChatHandler) HandleChatV1(c *gin.Context) {
	chatGptRequest, err := ch.parseRequest(c)
	if err != nil {
		c.JSON(constant.HTTPStatusCodeBadRequest, "Invalid request parameter")
		fmt.Println("error happed: ", err.Error())
		return
	}

	openAiKey := strings.TrimSpace(chatGptRequest.OpenAiKey)
	if len(openAiKey) == 0 {
		c.JSON(constant.HTTPStatusCodeBadRequest, "OpenAiKey is not provided")
		return
	}

	messages, respMessage, err := ch.sendChatRequest(chatGptRequest, c)
	if err != nil {
		fmt.Println("error happed", err.Error())
		c.JSON(constant.HTTPStatusCodeInternalError, ChatGptResponse{
			Content: "",
			Error:   err.Error(),
		})
	}

	context, err := ch.parseResponse(messages, respMessage, c)
	if err != nil {
		c.JSON(constant.HTTPStatusCodeInternalError, ChatGptResponse{
			Content: "",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(constant.HTTPStatusCodeSuccess, ChatGptResponse{
		Content: respMessage,
		Context: context,
	})
}

func (ch *ChatHandler) HandleChatV2(c *gin.Context) {
	chatGptRequest, err := ch.parseRequest(c)
	if err != nil {
		c.JSON(constant.HTTPStatusCodeBadRequest, "Invalid request parameter")
		fmt.Println("error happed: ", err.Error())
		return
	}

	messages, respMessage, err := ch.sendChatRequest(chatGptRequest, c)
	if err != nil {
		fmt.Println("error happed", err.Error())
		c.JSON(constant.HTTPStatusCodeInternalError, ChatGptResponse{
			Content: "",
			Error:   err.Error(),
		})
	}

	context, err := ch.parseResponse(messages, respMessage, c)
	if err != nil {
		c.JSON(constant.HTTPStatusCodeInternalError, ChatGptResponse{
			Content: "",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(constant.HTTPStatusCodeSuccess, ChatGptResponse{
		Content: respMessage,
		Context: context,
	})
}

func (*ChatHandler) parseResponse(messages []gogpt.ChatCompletionMessage, respMessage string, c *gin.Context) (string, error) {
	messages = append(messages, gogpt.ChatCompletionMessage{
		Role:    ROLE_AI,
		Content: respMessage,
	})

	context, err := json.Marshal(messages)
	if err != nil {
		return "", err
	}
	return string(context), nil
}

func (ch *ChatHandler) sendChatRequest(chatGptRequest ChatGptRequest, c *gin.Context) ([]gogpt.ChatCompletionMessage, string, error) {
	messages := ch.genChatMessages(chatGptRequest)

	respMessage, err := ch.chatService.Chat(c, nil, messages, chatGptRequest.Model, true)
	if err != nil {
		fmt.Println("error heppened", err.Error())
		return messages, "", err
	}
	return messages, respMessage, nil
}

func (*ChatHandler) genChatMessages(chatGptRequest ChatGptRequest) []gogpt.ChatCompletionMessage {
	messages := append(chatGptRequest.MsgHistory, gogpt.ChatCompletionMessage{
		Role:    ROLE_USER,
		Content: chatGptRequest.Message,
	})

	start := 0
	if len(messages) > MAX_CONTEXT_SIZE {
		start = len(messages) - MAX_CONTEXT_SIZE
	}

	messages = messages[start:]
	return messages
}

func (*ChatHandler) parseRequest(c *gin.Context) (ChatGptRequest, error) {
	var chatGptRequest ChatGptRequest
	err := c.BindJSON(&chatGptRequest)
	if err != nil {
		fmt.Println("error happed: ", err.Error())
	}
	fmt.Println("Received a request: ", chatGptRequest)

	model := c.Request.URL.Query().Get("model")
	if strings.TrimSpace(model) == "" {
		model = gogpt.GPT3Dot5Turbo
	}
	chatGptRequest.Model = model

	// var msgHistory []gogpt.ChatCompletionMessage
	fmt.Printf("current message: %v\n", chatGptRequest.Context)

	if len(chatGptRequest.Context) != 0 {
		err = json.Unmarshal([]byte(chatGptRequest.Context), &chatGptRequest.MsgHistory)
		if err != nil {
			fmt.Println("unmarshal chatgptRequest.Context error happed: ", err.Error())
		}

		fmt.Printf("current message: %v\n", chatGptRequest.Context)
	}

	return chatGptRequest, err
}
