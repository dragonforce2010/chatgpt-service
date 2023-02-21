package chat

type ChatGptResponse struct {
	Content        string `json:"content"`
	Context        string `json:"context"`
	Error          string `json:"error"`
}

type ChatGptRequest struct {
	Message        string `json:"message"`
	Context        string `json:"context"`
}
