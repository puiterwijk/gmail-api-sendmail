package main

import (
	"encoding/json"
	"log"
	"flag"
	"io/ioutil"
	"os"

	"github.com/puiterwijk/gmail-api-sendmail/internal"
)

// Ignored flags, to support sendmail API
var flags_to_ignore = []string{
	"Ac", "Am",
	"B", "ba", "bd", "bD", "bh","bH", "bi", "bm", "bp", "bP", "bs", "bt", "bv",
	"C",
	"D", "d",
	"F", "f",
	"L",
	"N", "n",
	"O", "o",
	"p",
	"q", "qp", "qf", "Q",
	"R", "r",
	"t",
	"V", "v",
	"X",
}

var (
	ignoreDots = flag.Bool("i", false, "Ignore dots on lines")
)

func init() {
	flaglist := make(map[string]*bool)

	for _, flagname := range flags_to_ignore {
		flaglist[flagname] = flag.Bool(flagname, false, "Ignored")
	}

	// Parse arguments
	flag.Parse()

	// Check various flags are not used
	for flagname, val := range flaglist {
		if *val {
			log.Fatalf("Flag %s is not supported", flagname)
		}
	}
}

func doqueue() bool {
	if os.Getenv("GMAIL_QUEUE_NEVER") != "" {
		return false
	}
	if os.Getenv("GMAIL_QUEUE_ALWAYS") != "" {
		return true
	}


	if _, err := os.Stat(internal.Queuedir()); err != nil {
		return false
	}
	return true
}

func main() {
	if flag.NArg() < 1 {
		log.Fatalf("Please provide at least one addressee")
	}

	var req internal.EmailRequest
	req.Destinations = flag.Args()

	msg, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading full message: %s", err)
	}
	req.MessageBody = string(msg)

	if doqueue() {
		qfile, err := ioutil.TempFile(internal.Queuedir(), "*.json")
		if err != nil {
			log.Fatalf("Error creating email queue file: %s", err)
		}
		err = json.NewEncoder(qfile).Encode(req)
		if err != nil {
			os.Remove(qfile.Name())
			log.Fatalf("Error queueing email: %s", err)
		}
		log.Println("Email succesfully queued")
		return
	}

	if err = internal.SendEmail(req); err != nil {
		log.Fatalf("Error sending email: %s", err)
	}
}
