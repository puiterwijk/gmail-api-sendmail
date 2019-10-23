package internal

import (
	"encoding/base64"
	"fmt"

	gmail "google.golang.org/api/gmail/v1"
)

func SendEmail(req EmailRequest) error {
	clt := getGmailClient()
	var msg gmail.Message

	msg.Raw = base64.RawURLEncoding.EncodeToString([]byte(req.MessageBody))

	_, err := clt.Users.Messages.Send("me", &msg).Do()
	if err != nil {
		return fmt.Errorf("Unable to send: %v", err)
	}

	return nil
}
