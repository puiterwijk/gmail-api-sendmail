package main

import (
	"encoding/json"
	"log"
	"io/ioutil"
	"path"
	"os"

	"github.com/puiterwijk/gmail-api-sendmail/internal"
)

func main() {
	files, err := ioutil.ReadDir(internal.Queuedir())
	if err != nil {
		log.Fatalf("Error getting list of emails to send: %s", err)
	}

	for _, f := range files {
		fname := path.Join(internal.Queuedir(), f.Name())
		f, err := os.Open(fname)
		if err != nil {
			log.Fatalf("Error opening %s: %s", fname, err)
		}
		defer f.Close()

		var req internal.EmailRequest
		err = json.NewDecoder(f).Decode(&req)
		if err != nil {
			log.Fatalf("Error parsing request %s: %s", fname, err)
		}

		if err = internal.SendEmail(req); err != nil {
			log.Fatalf("Error sending email for %s: %s", fname, err)
		}
		os.Remove(fname)
		log.Println("Sent email at ", fname)
	}
}
