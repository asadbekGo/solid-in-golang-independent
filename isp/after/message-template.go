package after

import (
	"fmt"
	"time"
)

type IMessageTemplate interface {
	Create() string
	CreateHtml() string
}

type MessageTemplate struct {
	StartedAt time.Time
	EndedAt   time.Time
}

type MessageTemplateWebinar struct {
	MessageTemplate
	Author string
	Title  string
}

type MessageTemplateCompetation struct {
	MessageTemplate
	Name  string
	Level string
}

func (m *MessageTemplateWebinar) New() *MessageTemplateWebinar {
	m = &MessageTemplateWebinar{
		MessageTemplate: MessageTemplate{
			StartedAt: time.Now(),
			EndedAt:   time.Now().Add(1 * time.Hour),
		},
		Author: "William Shekspir",
		Title:  "About Sherlock Holms",
	}

	return m
}

func (m *MessageTemplateWebinar) Create() string {
	v := m.New()
	res := fmt.Sprintf("Webinar Author: %s, Title: %s, Start time: %s, End time: %s", v.Author, v.Title, v.MessageTemplate.StartedAt, v.MessageTemplate.EndedAt)
	return res
}

func (m *MessageTemplateWebinar) CreateHtml() string {
	v := m.New()
	res := fmt.Sprintf("[Html] Webinar Author: %s, Title: %s, Start time: %s, End time: %s", v.Author, v.Title, v.MessageTemplate.StartedAt, v.MessageTemplate.EndedAt)
	return res
}

func (m *MessageTemplateCompetation) New() *MessageTemplateCompetation {
	m = &MessageTemplateCompetation{
		MessageTemplate: MessageTemplate{
			StartedAt: time.Now(),
			EndedAt:   time.Now().Add(1 * time.Hour),
		},
		Name:  "System Design",
		Level: "Hard",
	}

	return m
}

func (m *MessageTemplateCompetation) Create() string {
	v := m.New()
	res := fmt.Sprintf("Competation: %s, Level: %s, Start time: %s, End time: %s", v.Name, v.Level, v.MessageTemplate.StartedAt, v.MessageTemplate.EndedAt)
	return res
}

func (m *MessageTemplateCompetation) CreateHtml() string {
	v := m.New()
	res := fmt.Sprintf("[Html] Competation: %s, Level: %s, Start time: %s, End time: %s", v.Name, v.Level, v.MessageTemplate.StartedAt, v.MessageTemplate.EndedAt)
	return res
}
