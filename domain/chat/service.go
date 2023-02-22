package chat

import (
	"fmt"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/dragonforce2010/chatgpt-service/client"
	"github.com/gin-gonic/gin"
)

type ChatService struct {
	client *client.Client
}

func NewChatService(client *client.Client) *ChatService {
	return &ChatService{client: client}
}

func (c *ChatService) GetChatResponse(ctx *gin.Context, prompt []string) (string, error) {
	var maxToken = 3000
	resp, err := c.client.GptClient.CompletionWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt:    prompt,
		MaxTokens: &maxToken,
	})

	if err != nil || resp.Choices == nil {
		fmt.Println(err.Error())
		return "something wrong, not able to generate your response", err
	}

	fmt.Println("get result for chatgpt:", resp.Choices[0].Text)
	return resp.Choices[0].Text, nil
}

// func (c *ChatService) GetChatResponseV2(ctx *gin.Context, chatMessage string) (string, error) {
// 	var maxToken = 3000
// 	resp, err := gogpt.NewClient("").CreateCompletion(ctx, gogpt.CompletionRequest{

// 	})

// 	if err != nil || resp.Choices == nil {
// 		fmt.Println(err.Error())
// 		return "something wrong, not able to generate your response", err
// 	}

// 	fmt.Println("get result for chatgpt:", resp.Choices[0].Text)
// 	return resp.Choices[0].Text, nil
// }
