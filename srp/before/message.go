package before

import "fmt"

type IMessage interface {
	Create()
	Send()
}

type Message struct {
	Body     string
	Sender   int
	Receiver int
}

func (m *Message) Create() *Message {
	msg := Message{
		Body:     "Hello Mike",
		Sender:   1324,
		Receiver: 6031,
	}

	fmt.Println("Create Message")
	fmt.Printf("%v\n", msg)
	return &msg
}

func (m *Message) Send(msg *Message) *Message {
	fmt.Println("Send Message")
	fmt.Println("Sender:", msg.Sender)
	fmt.Println("Receiver:", msg.Receiver)
	return msg
}
