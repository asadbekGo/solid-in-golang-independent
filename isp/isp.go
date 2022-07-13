package isp

import (
	"fmt"

	"github.com/asadbekGo/solid-in-golang-independent/isp/after"
)

func Run() {
	fmt.Println("----->Run isp(Interface Segregation Principle)")

	messageTemplateWebinar := &after.MessageTemplateWebinar{}
	msg := &after.MessageSocmed{
		MessageTemplate: messageTemplateWebinar,
	}
	msgEmail := &after.MessageEmail{
		MessageTemplate: messageTemplateWebinar,
	}

	messageSocmedPayload := msg.Create()
	messageEmailPayload := msgEmail.Create()
	fmt.Println()

	user := &after.User{}

	senderSocmed := after.SenderSomed{
		Sender:   user.GetSender(),
		Receiver: user.GetReceiver(),
		Message:  messageSocmedPayload.Body,
	}

	senderEmail := after.SenderEmail{
		Sender:   user.GetSender(),
		Receiver: user.GetReceiver(),
		Subject:  messageEmailPayload.Subject,
		Body:     messageEmailPayload.Body,
	}

	senderSocmed.SendWhatsapp()
	fmt.Println()
	senderSocmed.SendTelegram()
	fmt.Println()
	senderEmail.SendEmail()
	fmt.Println()
}
