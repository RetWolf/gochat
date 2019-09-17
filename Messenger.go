package gochat

// Operator contains the base interface for a strategy pattern
type Operator interface {
	PrintMessage(e Event)
}

// Messenger is the generic interface which will allow us to change behavoir during runtime
type Messenger struct {
	Operator Operator
}

// SendMessage proxies the message to be sent to the Operator
func (m *Messenger) SendMessage(e Event) {
	m.Operator.PrintMessage(e)
}
