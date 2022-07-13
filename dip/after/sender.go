package after

import "fmt"

type ISenderSomed interface {
	SendWhatsapp()
	SendTelegram()
}

type SenderSomed struct {
	Sender   *User
	Receiver *User
	Message  string
}

type ISenderEmail interface {
	SendEmail()
}

type SenderEmail struct {
	Sender   *User
	Receiver *User
	Subject  string
	Body     string
}

func (s *SenderSomed) SendWhatsapp() bool {
	fmt.Println("Send Whatsapp")
	fmt.Println("Sender: ", s.Sender.Phone)
	fmt.Println("Receiver: ", s.Receiver.Phone)
	fmt.Println("Body: ", s.Message)
	return true
}

func (s *SenderSomed) SendTelegram() bool {
	fmt.Println("Send Telegram")
	fmt.Println("Sender: ", s.Sender.Phone)
	fmt.Println("Receiver: ", s.Receiver.Phone)
	fmt.Println("Body: ", s.Message)
	return true
}

func (s *SenderEmail) SendEmail() bool {
	fmt.Println("Send Telegram")
	fmt.Println("Sender: ", s.Sender.Phone)
	fmt.Println("Receiver: ", s.Receiver.Phone)
	fmt.Println("Subject: ", s.Subject)
	return true
}
