package mail_to_slack

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

var Slack *SlackHandler

func init() {
	token := os.Getenv("SLACK_BOT_TOKEN")
	Slack = NewSlackHandler(token)
}

func Start() {
	const maxLines = 30
	var channel = os.Getenv("SLACK_CHANNEL")

	b := new(bytes.Buffer)
	io.Copy(b, os.Stdin)

	fmt.Printf("%s", b)

	var mail = NewMailHandler(b.String())
	mail.Parse()

	var blockText SlackBlockText
	blockText.Text = mail.Subject
	blockText.Type = "plain_text"

	var header SlackBlock
	header.Type = "header"
	header.Text = blockText

	blockText = SlackBlockText{}
	blockText.Type = "mrkdwn"

	var sections = []SlackBlock{}
	var section SlackBlock
	section.Type = "section"

	for i, t := range strings.Split(mail.Message, "\n") {
		blockText.Text += "> " + t + "\n"

		if i > 0 && i%maxLines == 0 {
			blockText.Text = strings.TrimSuffix(blockText.Text, "\n")
			section.Text = blockText

			sections = append(sections, section)

			blockText.Text = ""
		}
	}

	blockText.Text = strings.TrimSuffix(blockText.Text, "\n")
	section.Text = blockText

	sections = append(sections, section)

	var blocks = append([]SlackBlock{header}, sections...)

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "UnknownServer"
	}

	err = Slack.PostMessage(SlackPostMessage{
		IconEmoji: os.Getenv("SLACK_ICON_EMOJI"),
		Channel:   channel,
		Blocks:    blocks,
		UserName:  fmt.Sprintf("[%s]Mail[%s]", strings.ToUpper(hostname), mail.From),
	})

	if err != nil {
		fmt.Println(err)
	}
}
