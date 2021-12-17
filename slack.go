package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

const SlackApiURI = "https://slack.com/api/"

type SlackHandler struct {
	token string
}

type SlackPostMessage struct {
	Token       string            `json:"token"`
	Channel     string            `json:"channel"`
	Attachments []SlackAttachment `json:"attachments,omitempty"`
	Blocks      []SlackBlock      `json:"blocks,omitempty"`
	UserName    string            `json:"username"`
	AsUser      bool              `json:"as_user"`
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
	OK    string `json:"ok"`
	Error string `json:"error"`
}

func NewSlackHandler(token string) *SlackHandler {
	return &SlackHandler{
		token: token,
	}
}

func (s SlackHandler) PostMessage(postMessage SlackPostMessage) (err error) {
	if postMessage.Token == "" {
		postMessage.Token = s.token
	}
	var buf = new(bytes.Buffer)
	err = json.NewEncoder(buf).Encode(postMessage)
	if err != nil {
		return
	}
	resp, err := http.Post(SlackApiURI, "application/json", buf)
	if err != nil {
		return
	}

	var response SlackResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return
	}

	if response.OK != "ok" {
		return errors.New(response.Error)
	}

	return nil
}
