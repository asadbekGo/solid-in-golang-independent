package ocp

import (
	"fmt"

	"github.com/asadbekGo/solid-in-golang-independent/ocp/after"
)

func Run() {
	fmt.Println("----->Run ocp(Open-Closed Principle)")

	messageTemplateWebinar := &after.MessageTemplateWebinar{}
	msg := &after.Message{
		MessageTemplateWebinar: messageTemplateWebinar,
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
