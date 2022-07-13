package after

import "fmt"

type IMessage interface {
	Create() *Message
}

type Message struct {
	Body                       string
	MessageTemplateWebinar     *MessageTemplateWebinar
	MessageTemplateCompetation *MessageTemplateCompetation
}

func (m *Message) Create() string {
	var template string
	if m.MessageTemplateWebinar != nil {
		template = m.MessageTemplateWebinar.Create()
	}

	if m.MessageTemplateCompetation != nil {
		template = m.MessageTemplateCompetation.Create()
	}

	m = &Message{
		Body: "Hi John ! " + template,
	}

	res := m.Body

	fmt.Println("Create message")
	fmt.Printf("%+v\n", m)

	return res
}
