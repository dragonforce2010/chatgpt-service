package client

import (
	"math/rand"
	"os"
	"strings"
	"sync"

	"github.com/dragonforce2010/chatgpt-service/constant"
	gpt3 "github.com/sashabaranov/go-gpt3"
)

var once sync.Once

type Client struct {
	GptClient   *gpt3.Client
	ClientsPool []*gpt3.Client
}

func NewClient() *Client {
	return (&Client{}).initClient()
}

func (c *Client) initClient() *Client {
	once.Do(func() {
		openaiAPIKeyString := os.Getenv(constant.OpenaiApiKey)

		if openaiAPIKeyString == "" {
			panic("Failed to get open ai api key")
		}

		for _, key := range strings.Split(openaiAPIKeyString, ",") {
			c.ClientsPool = append(c.ClientsPool, gpt3.NewClient(strings.TrimSpace(key)))
		}
	})
	return c
}

func (c *Client) GetRandomOneClient() *gpt3.Client {
	return c.ClientsPool[rand.Intn(len(c.ClientsPool))]
}
