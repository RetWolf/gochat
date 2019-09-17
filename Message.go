package gochat

// Message is a generic interface to contain all Typed messaged
type Message interface {
	Operator
	GetMessage() string
	GetMessageType() MessageType
}
