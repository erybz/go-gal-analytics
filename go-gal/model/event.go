package model

type Event struct {
	Location EventLocation `json:"location"`
	Device   EventDevice   `json:"device"`
}

type EventLocation struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

type EventDevice struct {
	Type     string `json:"type"`
	Platform string `json:"platform"`
	OS       string `json:"os"`
	Browser  string `json:"browser"`
}

func (e *Event) Valid() bool {
	if e.Location.City != "" ||
		e.Location.Country != "" ||
		e.Device.Type != "" ||
		e.Device.Platform != "" ||
		e.Device.OS != "" ||
		e.Device.Browser != "" {
		return true
	}
	return false
}
