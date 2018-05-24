package messaging

import (
	"context"

	"firebase.google.com/go/messaging"
	"golang.org/x/oauth2/google"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Handler messaging handler
type Handler struct {
	client *messaging.Client
}

// MessageLoc message with loc
type MessageLoc struct {
	TitleLocKey  string
	TitleLocArgs []string
	BodyLocKey   string
	BodyLocArgs  []string
	Sound        string
	Badge        int
	Priority     string
}

// Message message with title - body
type Message struct {
	Title    string
	Body     string
	Sound    string
	Badge    int
	Priority string
}

// InitWithCredentialsFile init messaging handler with credentials file
func InitWithCredentialsFile(credentials string) (Handler, error) {
	opt := option.WithCredentialsFile(credentials)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return Handler{}, err
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		return Handler{}, err
	}

	return Handler{
		client: client,
	}, nil
}

// InitWithCredentials init messaging handler with credentials object
func InitWithCredentials(creds *google.Credentials) (Handler, error) {
	opt := option.WithCredentials(creds)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return Handler{}, err
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		return Handler{}, err
	}

	return Handler{
		client: client,
	}, nil
}

// SendToTopicWithLoc send to topic with bodyLockKey and bodyLocArgs
func (h Handler) SendToTopicWithLoc(topic string, message MessageLoc) error {
	messageSend, err := h.GenerateMessageWithLoc(message)
	if err != nil {
		return err
	}
	messageSend.Topic = topic
	_, err = h.client.Send(context.Background(), &messageSend)
	return err
}

// SendToTokenWithLoc send to token with bodyLockKey and bodyLocArgs
func (h Handler) SendToTokenWithLoc(token string, message MessageLoc) error {
	messageSend, err := h.GenerateMessageWithLoc(message)
	if err != nil {
		return err
	}
	messageSend.Token = token
	_, err = h.client.Send(context.Background(), &messageSend)
	return err
}

// SendToTopic send to topic with title and body
func (h Handler) SendToTopic(topic string, message Message) error {
	messageSend, err := h.GenerateMessage(message)
	if err != nil {
		return err
	}
	messageSend.Topic = topic
	_, err = h.client.Send(context.Background(), &messageSend)
	return err
}

// SendToToken send to token with title and body
func (h Handler) SendToToken(token string, message Message) error {
	messageSend, err := h.GenerateMessage(message)
	if err != nil {
		return err
	}
	messageSend.Token = token
	_, err = h.client.Send(context.Background(), &messageSend)
	return err
}
