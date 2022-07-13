package dip

import (
	"fmt"

	"github.com/asadbekGo/solid-in-golang-independent/dip/after"
)

type Factory struct {
	MessageSocmed after.IMessageSocmed
	MessageEmail  after.IMessageEmail
	User          after.IUser
}

func Run() {
	fmt.Println("----->Run dip(Dependency Inversion Principle)")

	messageTemplateWebinar := &after.MessageTemplateWebinar{}
	MessageSocmed := &after.MessageSocmed{
		MessageTemplate: messageTemplateWebinar,
	}
	MessageEmail := &after.MessageEmail{
		MessageTemplate: messageTemplateWebinar,
	}

	// messageSocmedPayload := msg.Create()
	// messageEmailPayload := msgEmail.Create()
	// fmt.Println()

	User := &after.User{}

	factory := Factory{
		MessageSocmed: MessageSocmed,
		MessageEmail:  MessageEmail,
		User:          User,
	}

	Execute(&factory)
}

func Execute(f *Factory) {

	senderSocmed := after.SenderSomed{
		Sender:   f.User.GetSender(),
		Receiver: f.User.GetReceiver(),
		Message:  f.MessageSocmed.Create().Body,
	}

	senderEmail := after.SenderEmail{
		Sender:   f.User.GetSender(),
		Receiver: f.User.GetReceiver(),
		Subject:  f.MessageEmail.Create().Subject,
		Body:     f.MessageEmail.Create().Body,
	}

	senderSocmed.SendWhatsapp()
	fmt.Println()
	senderSocmed.SendTelegram()
	fmt.Println()
	senderEmail.SendEmail()
	fmt.Println()
}
