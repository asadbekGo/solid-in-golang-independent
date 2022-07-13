package after

import (
	"fmt"
	"time"
)

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

func (m *MessageTemplateWebinar) Create() string {
	m = &MessageTemplateWebinar{
		MessageTemplate: MessageTemplate{
			StartedAt: time.Now(),
			EndedAt:   time.Now().Add(1 * time.Hour),
		},
		Author: "William Shekspir",
		Title:  "About Sherlock Holms",
	}

	res := fmt.Sprintf(
		"Webinar Author: %s, Title: %s, Start time: %s, End time: %s",
		m.Author, m.Title, m.MessageTemplate.StartedAt, m.MessageTemplate.EndedAt,
	)

	return res
}

func (m *MessageTemplateCompetation) Create() string {
	m = &MessageTemplateCompetation{
		MessageTemplate: MessageTemplate{
			StartedAt: time.Now(),
			EndedAt:   time.Now().Add(1 * time.Hour),
		},
		Name:  "System Design",
		Level: "Hard",
	}

	res := fmt.Sprintf(
		"Competation: %s, Level: %s, Start time: %s, End time: %s",
		m.Name, m.Level, m.MessageTemplate.StartedAt, m.MessageTemplate.EndedAt,
	)

	return res
}
