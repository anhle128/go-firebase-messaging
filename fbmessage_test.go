package messaging_test

import (
	"testing"

	messaging "github.com/anhle128/go-firebase-messaging"
)

func TestSendMessage(t *testing.T) {

	notificationProcessor, err := messaging.InitWithCredentialsFile("./footballx-dev-firebase-adminsdk-09j5o-7caf574abf.json")
	if err != nil {
		t.Error(err)
		return
	}

	data := make(map[string]string)
	data["name"] = "le duc anh"
	data["age"] = "18"

	err = notificationProcessor.SendToToken("eCSSsvUhtSU:APA91bEdQXMbZh-Qz6RQyKBfTOFZG3aZptSeknIq5NRBqUfk_K0ZzRzxvOaI5jg7L5uDxiHRHk9Ik09rQ0tZkwhrf1U5DEqO0F0qU5XNDF2hj_xBmcV0URu57tdiPvKcjbqTgJrL6nWjDuhj3QFhTL4Y9Nt8RKrIAA", messaging.Message{
		Title: "Le Duc Anh",
		Body:  "Dep trai vl luon",
		Data:  data,
	})
	if err != nil {
		t.Error(err)
	}
}
