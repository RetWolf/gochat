package gochat

// MessageType holds the unique identifier for the created type
// Not meant to be used directly
type MessageType uint

const (
	// TEXT represents a message containing text
	TEXT MessageType = iota
	// IMAGE represents a message containing an image
	IMAGE MessageType = iota
	// GIF represents a message containing a gif
	GIF MessageType = iota
	// STICKER represents a message containing a sticker
	STICKER MessageType = iota
)
