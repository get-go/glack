package glack

import "github.com/nlopes/slack"

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
	Send(message *Message) (id string, err error)
}

//Send a simple message to the specified channel
func (c *Client) Send(message *Message) (id string, err error) {
	p := slack.PostMessageParameters{
		AsUser:    false,
		Username:  message.Username,
		IconEmoji: message.Icon,
	}

	id, _, err = c.client.PostMessage(message.Channel, message.Message, p)

	return id, err
}

//New creates a new Client object
func New(t string) Client {
	return Client{
		Token:  t,
		client: slack.New(t),
	}
}
