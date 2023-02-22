package chat

import (
	"fmt"

	"github.com/dragonforce2010/chatgpt-service/client"
	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
)

type ChatService struct {
	client *client.Client
}

func NewChatService(client *client.Client) *ChatService {
	return &ChatService{client: client}
}

func (c *ChatService) GetChatResponse(ctx *gin.Context, prompt string) (string, error) {
	var maxToken int = 3000
	client := c.client.GetRandomOneClient()
	resp, err := client.CreateCompletion(ctx, gogpt.CompletionRequest{
		Prompt:    prompt,
		Suffix:    "",
		MaxTokens: maxToken,
		Model:     gogpt.GPT3Ada,
	})

	if err != nil || resp.Choices == nil {
		fmt.Println(err.Error())
		return "something wrong, not able to generate your response", err
	}

	fmt.Println("get result for chatgpt:", resp.Choices[0].Text)
	return resp.Choices[0].Text, nil
}
