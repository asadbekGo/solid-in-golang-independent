package srp

import (
	"fmt"

	"github.com/asadbekGo/solid-in-golang-independent/srp/after"
)

func Run() {
	fmt.Println("----->Run srp(Single Responsibility Principle)")

	msg := &after.Message{}
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
