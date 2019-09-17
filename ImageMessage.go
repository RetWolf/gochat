package gochat

import "fmt"

// ImageMessage holds the type and message
// Do not create using this struct, use the NewImageMessage constructor
type ImageMessage struct {
	messageType MessageType
	message     string
}

// SetMessage is used to change the message at runtime
func (m *ImageMessage) SetMessage(message string) {
	m.message = message
}

// PrintMessage prints the message to the terminal
func (m *ImageMessage) PrintMessage(e Event) {
	e.TargetRoom.OnNotify(e)
	fmt.Printf("Message from %s to Room %s: %s\n", e.Sender.Name, e.TargetRoom.ID, m.message)
}

// GetMessageType returns the message type for this particular message
func (m *ImageMessage) GetMessageType() MessageType {
	return m.messageType
}

// GetMessage returns the current message
func (m *ImageMessage) GetMessage() string {
	return m.message
}

// NewImageMessage creates a new image message with a constant message
func NewImageMessage() ImageMessage {
	return ImageMessage{
		messageType: IMAGE,
		message:     "This is an image, open me in Photoshop!",
	}
}
