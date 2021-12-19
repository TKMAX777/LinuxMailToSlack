package mail_to_slack

import (
	"strings"
	"time"
)

type MailHandler struct {
	Text                    string
	From                    string
	To                      string
	Date                    time.Time
	Subject                 string
	ContentType             string
	ContentTransferEncoding string
	MINEversion             string
	Message                 string
}

func NewMailHandler(text string) *MailHandler {
	return &MailHandler{Text: text}
}

func (m *MailHandler) Parse() {
	var messages = strings.Split(m.Text, "\n")

MessageLoop:
	for i, s := range messages {
		attrs := strings.Split(s, ": ")

		switch strings.TrimSpace(attrs[0]) {
		case "Delivery-date":
			t, err := time.Parse(time.RFC1123Z, attrs[1])
			if err != nil {
				continue
			}
			m.Date = t
		case "From":
			m.From = attrs[1]
		case "To":
			m.To = attrs[1]
		case "Subject":
			m.Subject = attrs[1]
		case "MIME-Version":
			m.MINEversion = attrs[1]
		case "Content-Type":
			m.ContentType = attrs[1]
		case "Content-Transfer-Encoding":
			m.ContentTransferEncoding = attrs[1]
		case "":
			m.Message = strings.Join(messages[i+1:], "\n")
			break MessageLoop
		}
	}
}
