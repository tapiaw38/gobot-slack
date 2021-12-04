package message

import (
	"bot-slack/models"
)

// Func sendMessage - send message to slack
func SendMessage(name string, message string) string {

	m := models.Message{
		Name:    name,
		Message: message,
	}

	return "name: " + m.Name + " message: " + m.Message
}
