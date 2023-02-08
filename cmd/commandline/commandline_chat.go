package commandline

import (
	"bufio"
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/cobra"
	"os"
	"sync"
)

const API_KEY = "CHATGPT_API_KEY"

var client gpt3.Client

func initClient() {
	var once sync.Once
	once.Do(func() {
		apiKey := os.Getenv(API_KEY)
		if apiKey == "" {
			apiKey = "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"
			//panic("Missing API KEY!")
		}

		client = gpt3.NewClient(apiKey)
	})
}

func GetChatResponse(ctx context.Context, chatMessage string) {

	var temp float32 = 0.9
	var maxToken = 512
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt:      []string{chatMessage},
		MaxTokens:   &maxToken,
		Temperature: &temp,
	}, func(resp *gpt3.CompletionResponse) {
		fmt.Print(resp.Choices[0].Text)
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
}

func main2() {
	initClient()
	ctx := context.Background()
	rootCmd := &cobra.Command{
		Use:   "chatgpt",
		Short: "Chat with chatgpt",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			for !quit {
				fmt.Println("Your message: ")

				if !scanner.Scan() {
					break
				}

				message := scanner.Text()
				switch message {
				case "quit":
					quit = true
				default:
					GetChatResponse(ctx, message)
				}
			}
		},
	}
	err := rootCmd.Execute()

	if err != nil {
		return
	}
}
