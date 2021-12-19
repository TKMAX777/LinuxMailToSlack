package main

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

func main() {
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

	var text string
	for _, t := range strings.Split(mail.Message, "\n") {
		text += "> " + t + "\n"
	}
	text = strings.TrimSuffix(text, "\n")

	blockText.Type = "mrkdwn"
	blockText.Text = text

	var section SlackBlock
	section.Type = "section"
	section.Text = blockText

	var blocks = []SlackBlock{}
	blocks = append(blocks, header, section)

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
