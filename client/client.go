package client

import (
	"github.com/PullRequestInc/go-gpt3"
	"github.com/dragonforce2010/chatgpt-service/constant"
	"os"
	"sync"
)

var once sync.Once

type Client struct {
	GptClient gpt3.Client
	apiKey    string
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
