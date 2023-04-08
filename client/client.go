package client

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
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
		openaiKeyCount, err := strconv.ParseInt(os.Getenv(constant.OPENAPI_PAI_KEY_COUNT), 10, 32)

		if err != nil {
			panic(err)
		}

		for i := 1; i <= int(openaiKeyCount); i++ {
			openAiKey := os.Getenv(constant.OPEN_AI_API_KEY_PREFIX + strconv.Itoa(i))

			if strings.TrimSpace(openAiKey) == "" {
				panic("OpenAiKey is not valid!")
			}
			c.ClientsPool = append(c.ClientsPool, gpt3.NewClient(openAiKey))
			fmt.Println("successfully initialized one gpt client with key", openAiKey[:8])
		}
	})
	return c
}

func (c *Client) GetRandomOneClient() *gpt3.Client {
	return c.ClientsPool[rand.Intn(len(c.ClientsPool))]
}

func (c *Client) GetUserClient(openAiKey string) *gpt3.Client {
	return gpt3.NewClient(openAiKey)
}
