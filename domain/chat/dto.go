package chat

type Message struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type ChatGptResponse struct {
	ID             string `json:"id"`
	ResponseID     string `json:"response_id"`
	ConversationID string `json:"conversation_id"`
	Content        string `json:"content"`
	Error          string `json:"error"`
}

type ChatGptRequest struct {
	MessageID      string `json:"message_id"`
	ConversationID string `json:"conversation_id"`
	ParentID       string `json:"parent_id"`
	Message        string `json:"message"`
}
