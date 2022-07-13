package after

import "fmt"

type IMessage interface {
	Create() *Message
}

type Message struct {
	Body            string
	MessageTemplate IMessageTemplate
}

func (m *Message) Create() string {
	template := m.MessageTemplate.Create()

	m = &Message{
		Body: "Hi John ! " + template,
	}

	res := m.Body

	fmt.Println("Create message")
	fmt.Printf("%+v\n", m)

	return res
}
