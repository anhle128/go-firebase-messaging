package messaging

import (
	"fmt"
	"strings"

	fbmessaging "firebase.google.com/go/messaging"
)

func (h Handler) generateMessageWithBodyLoc(message MessageLoc) fbmessaging.Message {
	return fbmessaging.Message{
		// notify data for andorid
		Android: &fbmessaging.AndroidConfig{
			CollapseKey: message.BodyLocKey,
			Priority:    message.Priority,
		},
		Data: map[string]string{
			"body_loc_key":  message.BodyLocKey,
			"body_loc_args": strings.Replace(fmt.Sprint(message.BodyLocArgs), " ", ",", -1),
			"sound":         message.Sound,
		},
		// notify data for ios
		APNS: &fbmessaging.APNSConfig{
			Payload: &fbmessaging.APNSPayload{
				Aps: &fbmessaging.Aps{
					Sound: message.Sound,
					Badge: &message.Badge,
					Alert: &fbmessaging.ApsAlert{
						LocKey:  message.BodyLocKey,
						LocArgs: message.BodyLocArgs,
					},
				},
			},
		},
		// notify data for web
		Webpush: &fbmessaging.WebpushConfig{
			Data: map[string]string{
				"body_loc_key":  message.BodyLocKey,
				"body_loc_args": strings.Replace(fmt.Sprint(message.BodyLocArgs), " ", ",", -1),
				"sound":         message.Sound,
			},
		},
		Notification: &fbmessaging.Notification{},
	}
}

func (h Handler) generateMessageWithTitleBodyLoc(message MessageLoc) fbmessaging.Message {
	return fbmessaging.Message{
		// notify data for andorid
		Android: &fbmessaging.AndroidConfig{
			CollapseKey: message.BodyLocKey,
			Priority:    message.Priority,
		},
		Data: map[string]string{
			"title_loc_key":  message.TitleLocKey,
			"title_loc_args": strings.Replace(fmt.Sprint(message.TitleLocArgs), " ", ",", -1),
			"body_loc_key":   message.BodyLocKey,
			"body_loc_args":  strings.Replace(fmt.Sprint(message.BodyLocArgs), " ", ",", -1),
			"sound":          message.Sound,
		},
		// notify data for ios
		APNS: &fbmessaging.APNSConfig{
			Payload: &fbmessaging.APNSPayload{
				Aps: &fbmessaging.Aps{
					Sound: message.Sound,
					Badge: &message.Badge,
					Alert: &fbmessaging.ApsAlert{
						LocKey:       message.BodyLocKey,
						LocArgs:      message.BodyLocArgs,
						TitleLocKey:  message.TitleLocKey,
						TitleLocArgs: message.TitleLocArgs,
					},
				},
			},
		},
		// notify data for web
		Webpush: &fbmessaging.WebpushConfig{
			Data: map[string]string{
				"title_loc_key":  message.TitleLocKey,
				"title_loc_args": strings.Replace(fmt.Sprint(message.TitleLocArgs), " ", ",", -1),
				"body_loc_key":   message.BodyLocKey,
				"body_loc_args":  strings.Replace(fmt.Sprint(message.BodyLocArgs), " ", ",", -1),
				"sound":          message.Sound,
			},
		},
	}
}

func (h Handler) generateMessageTitleBody(message Message) fbmessaging.Message {
	return fbmessaging.Message{
		Android: &fbmessaging.AndroidConfig{
			Priority: "high",
		},
		Data: map[string]string{
			"title": message.Title,
			"body":  message.Body,
			"sound": message.Sound,
		},
		APNS: &fbmessaging.APNSConfig{
			Payload: &fbmessaging.APNSPayload{
				Aps: &fbmessaging.Aps{
					Sound: message.Sound,
					Badge: &message.Badge,
					Alert: &fbmessaging.ApsAlert{
						Body:  message.Body,
						Title: message.Title,
					},
				},
			},
		},
		Webpush: &fbmessaging.WebpushConfig{
			Notification: &fbmessaging.WebpushNotification{
				Icon:  "https://my-server/icon.png",
				Title: message.Title,
				Body:  message.Body,
			},
		},
		Notification: &fbmessaging.Notification{},
	}
}

func (h Handler) generateMessageBody(message Message) fbmessaging.Message {
	return fbmessaging.Message{
		Android: &fbmessaging.AndroidConfig{
			Priority: "high",
		},
		Data: map[string]string{
			"body":  message.Body,
			"sound": message.Sound,
		},
		APNS: &fbmessaging.APNSConfig{
			Payload: &fbmessaging.APNSPayload{
				Aps: &fbmessaging.Aps{
					Sound: message.Sound,
					Badge: &message.Badge,
					Alert: &fbmessaging.ApsAlert{
						Body: message.Body,
					},
				},
			},
		},
		Webpush: &fbmessaging.WebpushConfig{
			Notification: &fbmessaging.WebpushNotification{
				Body: message.Body,
			},
		},
		Notification: &fbmessaging.Notification{},
	}
}

// GenerateMessageWithLoc generate message with loc
func (h Handler) GenerateMessageWithLoc(message MessageLoc) (fbmessaging.Message, error) {
	// default value
	if len(message.Sound) == 0 {
		message.Sound = "default"
	}
	// default value
	if message.Badge == 0 {
		message.Badge = 1
	}
	// default value
	if len(message.Priority) == 0 {
		message.Priority = "high"
	}
	if len(message.TitleLocKey) == 0 {
		return h.generateMessageWithBodyLoc(message), nil
	}
	if len(message.TitleLocArgs) == 0 {
		return fbmessaging.Message{}, fmt.Errorf("message.TitleLocArgs must have value")
	}
	return h.generateMessageWithTitleBodyLoc(message), nil
}

// GenerateMessage generate message with loc
func (h Handler) GenerateMessage(message Message) (fbmessaging.Message, error) {
	// default value
	if len(message.Sound) == 0 {
		message.Sound = "default"
	}
	// default value
	if message.Badge == 0 {
		message.Badge = 1
	}
	// default value
	if len(message.Priority) == 0 {
		message.Priority = "high"
	}
	if len(message.Title) == 0 {
		return h.generateMessageBody(message), nil
	}
	return h.generateMessageTitleBody(message), nil
}
