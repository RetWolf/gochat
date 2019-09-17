package gochat

import "fmt"

// GifMessage holds the type and message
// Do not create using this struct, use the NewGifMessage constructor
type GifMessage struct {
	messageType MessageType
	message     string
}

// SetMessage is used to change the message at runtime
func (m *GifMessage) SetMessage(message string) {
	m.message = message
}

// PrintMessage prints the message to the terminal
func (m *GifMessage) PrintMessage(e Event) {
	e.TargetRoom.OnNotify(e)
	fmt.Printf("Message from %s to Room %s: %s\n", e.Sender.Name, e.TargetRoom.ID, m.message)
}

// GetMessageType returns the message type for this particular message
func (m *GifMessage) GetMessageType() MessageType {
	return m.messageType
}

// GetMessage returns the current message
func (m *GifMessage) GetMessage() string {
	return m.message
}

// NewGifMessage creates a new gif message with a constant message
func NewGifMessage() GifMessage {
	return GifMessage{
		messageType: GIF,
		message:     "I'm a gif! Look at my high framerate hehe!",
	}
}
