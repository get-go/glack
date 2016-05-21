package glack

import (
	"strings"

	"github.com/nlopes/slack"
)

// Version of the glack SDK
var Version = "1.0.6"

// Message to send to slack
type Message struct {
	Channel  string
	Username string
	Message  string
	Icon     string
}

//Client has the token and slack client object
type Client struct {
	Token  string
	client *slack.Client
}

//Sender sends messages to the slack API
type Sender interface {
	Send(message *Message) (channelID, messageID string, err error)
}

//Send a simple message to the specified channel
func (c *Client) Send(message *Message) (channelID, messageID string, err error) {

	p := slack.PostMessageParameters{
		AsUser:   false,
		Username: message.Username,
	}

	if strings.HasPrefix(message.Icon, ":") && strings.HasSuffix(message.Icon, ":") {
		p.IconEmoji = message.Icon
	} else {
		p.IconURL = message.Icon
	}

	channelID, messageID, err = c.client.PostMessage(message.Channel, message.Message, p)

	return
}

//UploadFile to slack
func (c *Client) UploadFile(channel, filename string) (file *slack.File, err error) {
	fileUploadParameters := slack.FileUploadParameters{
		File:     filename,
		Channels: []string{channel},
	}

	return c.client.UploadFile(fileUploadParameters)
}

//New creates a new Client object
func New(t string) Client {
	return Client{
		Token:  t,
		client: slack.New(t),
	}
}
