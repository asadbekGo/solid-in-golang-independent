package lsp

import (
	"fmt"

	"github.com/asadbekGo/solid-in-golang-independent/lsp/after"
)

func Run() {
	fmt.Println("----->Run lsp(Liskov Subsituation Principle)")

	messageTemplateWebinar := &after.MessageTemplateWebinar{}
	msg := &after.Message{
		MessageTemplate: messageTemplateWebinar,
	}

	messagePayload := msg.Create()
	fmt.Println()

	user := &after.User{}

	sender := after.Sender{
		Sender:   user.GetSender(),
		Receiver: user.GetReceiver(),
		Message:  messagePayload,
	}

	sender.SendWhatsapp()
	fmt.Println()
	sender.SendTelegram()
	fmt.Println()
}
