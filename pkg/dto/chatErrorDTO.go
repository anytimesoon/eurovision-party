package dto

type ChatErrorDTO struct {
	Message string `json:"message"`
}

func NewChatErrorDTO(message string) ChatErrorDTO {
	return ChatErrorDTO{Message: message}
}
