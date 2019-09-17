package gochat

import "fmt"

// StickerMessage holds the type and message
// Do not create using this struct, use the NewStickerMessage constructor
type StickerMessage struct {
	messageType MessageType
	message     string
}

// SetMessage is used to change the message at runtime
func (m *StickerMessage) SetMessage(message string) {
	m.message = message
}

// PrintMessage prints the message to the terminal
func (m *StickerMessage) PrintMessage(e Event) {
	e.TargetRoom.OnNotify(e)
	fmt.Printf("Message from %s to Room %s: %s\n", e.Sender.Name, e.TargetRoom.ID, m.message)
}

// GetMessageType returns the message type for this particular message
func (m *StickerMessage) GetMessageType() MessageType {
	return m.messageType
}

// GetMessage returns the current message
func (m *StickerMessage) GetMessage() string {
	return m.message
}

// NewStickerMessage creates a new sticker message with a constant message
func NewStickerMessage() StickerMessage {
	return StickerMessage{
		messageType: STICKER,
		message:     "Microsoft Teams stickers are awesome",
	}
}
