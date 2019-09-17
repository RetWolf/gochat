package gochat

import (
	"fmt"
)

// TextMessage holds the type and message
// Do not create using this struct, use the NewTextMessage constructor
type TextMessage struct {
	messageType MessageType
	message     string
}

// SetMessage is used to change the message at runtime
func (m *TextMessage) SetMessage(message string) {
	m.message = message
}

// PrintMessage prints the text message to the terminal
func (m *TextMessage) PrintMessage(e Event) {
	e.TargetRoom.OnNotify(e)
	fmt.Printf("Message from %s to Room %s: %s\n", e.Sender.Name, e.TargetRoom.ID, m.message)
}

// GetMessageType returns the message type for this particular message
func (m *TextMessage) GetMessageType() MessageType {
	return m.messageType
}

// GetMessage returns the current message
func (m *TextMessage) GetMessage() string {
	return m.message
}

// NewTextMessage creates a new text message with a given message
func NewTextMessage(m string) TextMessage {
	return TextMessage{
		messageType: TEXT,
		message:     m,
	}
}
