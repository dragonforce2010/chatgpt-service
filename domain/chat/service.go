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

func (c *ChatService) GetChatResponse(ctx *gin.Context, prompt string, model string) (string, error) {
	maxToken, temperature, presencePenalty, frequencyPenalty := c.initParams(model)

	client := c.client.GetRandomOneClient()
	resp, err := client.CreateCompletion(ctx, gogpt.CompletionRequest{
		Prompt:           prompt,
		Suffix:           "",
		MaxTokens:        maxToken,
		Model:            model,
		Stop:             []string{prompt_context_prefix, prompt_question_prefix, prompt_question_postfix},
		Temperature:      temperature,
		PresencePenalty:  presencePenalty,
		FrequencyPenalty: frequencyPenalty,
		// Model:     gogpt.GPT3TextDavinci003,
	})

	if err != nil || resp.Choices == nil {
		fmt.Println(err.Error())
		return "something wrong, not able to generate your response", err
	}

	fmt.Println("get result for chatgpt:", resp.Choices[0].Text)
	return resp.Choices[0].Text, nil
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
