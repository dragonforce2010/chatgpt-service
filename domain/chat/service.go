package chat

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dragonforce2010/chatgpt-service/client"
	"github.com/dragonforce2010/chatgpt-service/constant"
	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
)

type ChatService struct {
	client *client.Client
}

func NewChatService(client *client.Client) *ChatService {
	return &ChatService{client: client}
}

func (c *ChatService) Chat(ctx *gin.Context, client *gogpt.Client, messages []gogpt.ChatCompletionMessage, model string, useClientPool bool) (string, error) {
	maxToken, temperature, presencePenalty, frequencyPenalty := c.initParams(model)

	if !useClientPool && client == nil {
		return "", fmt.Errorf("GptClient is nil")
	}

	if useClientPool {
		client = c.client.GetRandomOneClient()
	}
	return c.getChatResponse(client, ctx, messages, maxToken, model, temperature, presencePenalty, frequencyPenalty)
}

func (*ChatService) getChatResponse(client *gogpt.Client, ctx *gin.Context, messages []gogpt.ChatCompletionMessage, maxToken int, model string, temperature float32, presencePenalty float32, frequencyPenalty float32) (string, error) {
	fmt.Printf("model parameters - modelName: %v, maxToken: %v, temperature: %v, presencePenalty: %v, frequencyPenalty: %v\n", model, maxToken, temperature, presencePenalty, frequencyPenalty)
	resp, err := client.CreateChatCompletion(ctx, gogpt.ChatCompletionRequest{
		Model:            model,
		Messages:         messages,
		MaxTokens:        maxToken,
		Temperature:      temperature,
		Stream:           false,
		Stop:             []string{},
		PresencePenalty:  presencePenalty,
		FrequencyPenalty: frequencyPenalty,
	})

	if err != nil || resp.Choices == nil {
		fmt.Println(err.Error())
		return "something wrong, not able to generate your response", err
	}

	fmt.Println("get result for chatgpt:", resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content, nil
}

func (*ChatService) initParams(model string) (int, float32, float32, float32) {
	noneGpt3MaxToken, err := strconv.ParseInt(os.Getenv(constant.CHATGPT_PARAM_MAXTOKEN_NONE_GPT3), 10, 32)
	if err != nil {
		noneGpt3MaxToken = 1500
	}
	gpt3MaxToken, err := strconv.ParseInt(os.Getenv(constant.CHATGPT_PARAM_MAXTOKEN_GPT3), 10, 32)
	if err != nil {
		gpt3MaxToken = 3000
	}

	var maxToken = noneGpt3MaxToken
	if model == gogpt.GPT3TextDavinci003 {
		maxToken = gpt3MaxToken
	}

	temperature, err := strconv.ParseFloat(os.Getenv(constant.CHATGPT_PARAM_TEMPERATURE), 32)
	if err != nil {
		temperature = 0.7
	}

	presencePenalty, err := strconv.ParseFloat(os.Getenv(constant.CHATGPT_PARAM_PRESENCE_PENALTY), 32)
	if err != nil {
		presencePenalty = float64(0.6)
	}

	frequencyPenalty, err := strconv.ParseFloat(os.Getenv(constant.CHATGPT_PARAM_FREQUENCY_PENALTY), 32)
	if err != nil {
		frequencyPenalty = float64(0)
	}
	return int(maxToken), float32(temperature), float32(presencePenalty), float32(frequencyPenalty)
}
