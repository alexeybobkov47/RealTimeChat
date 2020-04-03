package structs

// MessageType -
type MessageType string

const (
	// MTPing -
	MTPing MessageType = "ping"
	// MTPong -
	MTPong MessageType = "pong"
	// MTMessage -
	MTMessage MessageType = "message"
)

// Message -
type Message struct {
	Type MessageType `json:"type"`
	Data string      `json:"data,omitempty"`
}

// Conf -
type Conf struct {
	LogFile string `json:"log_file" check:"required"`
	Port    string `json:"port" check:"required"`
}
