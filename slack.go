package mail_to_slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

const SlackApiURI = "https://slack.com/api"

type SlackHandler struct {
	token string
}

type SlackPostMessage struct {
	Channel     string            `json:"channel"`
	Attachments []SlackAttachment `json:"attachments,omitempty"`
	Blocks      []SlackBlock      `json:"blocks,omitempty"`
	UserName    string            `json:"username"`
	IconEmoji   string            `json:"icon_emoji"`
}

type SlackAttachment struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

type SlackBlock struct {
	Type string         `json:"type"`
	Text SlackBlockText `json:"text"`
}
type SlackBlockText struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji,omitempty"`
}

type SlackResponse struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
}

func NewSlackHandler(token string) *SlackHandler {
	return &SlackHandler{
		token: token,
	}
}

func (s SlackHandler) PostMessage(postMessage SlackPostMessage) (err error) {
	var buf = new(bytes.Buffer)
	err = json.NewEncoder(buf).Encode(postMessage)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", SlackApiURI+"/chat.postMessage", buf)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", "Bearer "+s.token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	var response SlackResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return
	}

	if !response.OK {
		return errors.New(response.Error)
	}

	return nil
}
