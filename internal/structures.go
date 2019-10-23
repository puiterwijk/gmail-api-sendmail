package internal

type EmailRequest struct {
	Flags struct {
		IgnoreDots bool `json:"ignoredots"`
	} `json:"flags"`

	Destinations []string `json:"destinations"`

	// Message
	MessageBody string `json:"messagebody"`
}
