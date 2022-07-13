package after

import "fmt"

type Sender struct {
	Sender   *User
	Receiver *User
	Message  string
}

func (s *Sender) SendWhatsapp() bool {
	fmt.Println("Send Whatsapp")
	fmt.Println("Sender: ", s.Sender.Phone)
	fmt.Println("Receiver: ", s.Receiver.Phone)
	fmt.Println("Body: ", s.Message)
	return true
}

func (s *Sender) SendTelegram() bool {
	fmt.Println("Send Telegram")
	fmt.Println("Sender: ", s.Sender.Phone)
	fmt.Println("Receiver: ", s.Receiver.Phone)
	fmt.Println("Body: ", s.Message)
	return true
}
