package internal

import (
	"log"
	"path"
	"os/user"
)

func Queuedir() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf("Error getting user: %s", err)
	}
	return path.Join(user.HomeDir, ".gmail-queue")
}
