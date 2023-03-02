package chat

import gogpt "github.com/sashabaranov/go-gpt3"

type ChatGptResponse struct {
	Content string `json:"content"`
	Context string `json:"context"`
	Error   string `json:"error"`
}

type ChatGptRequest struct {
	Message    string                        `json:"message,omitempty"`
	Context    string                        `json:"context,omitempty"`
	Model      string                        `json:"model,omitempty"`
	OpenAiKey  string                        `json:"openAiKey,omitempty"`
	MsgHistory []gogpt.ChatCompletionMessage `json:"-"`
}
