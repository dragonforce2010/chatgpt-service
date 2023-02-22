package client

import (
	"os"
	"sync"

	"github.com/dragonforce2010/chatgpt-service/constant"
	gpt3 "github.com/sashabaranov/go-gpt3"
)

var once sync.Once

type Client struct {
	GptClient *gpt3.Client
}

func NewClient() *Client {
	return (&Client{}).initClient()
}

func (c *Client) initClient() *Client {
	once.Do(func() {
		openaiAPIKey := os.Getenv(constant.OpenaiApiKey)
		if openaiAPIKey == "" {
			panic("Failed to get open ai api key")
		}
		c.GptClient = gpt3.NewClient(openaiAPIKey)
	})
	return c
}
