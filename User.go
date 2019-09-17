package gochat

import (
	"fmt"
)

// User contains a name and composes a messenger into our object
// Create using the NewUser constructor
type User struct {
	messenger *Messenger
	Name      string
	rooms     []Room
}

// SendMessage uses the strategy pattern to send any type of message
func (u *User) SendMessage(message Message, room Room) {
	e := Event{
		Sender:     *u,
		Message:    message,
		TargetRoom: room,
	}
	u.messenger = &Messenger{Operator: message}
	u.messenger.SendMessage(e)
}

// Register adds a room to a user. Used for bi-directional messaging
// Do not use these directly. Instead, register and deregister a user through the Room
func (u *User) Register(r Room) {
	u.rooms = append(u.rooms, r)
}

// Deregister remove a room from a user
func (u *User) Deregister(r Room) {
	for i, room := range u.rooms {
		if room.ID == r.ID {
			u.rooms[i] = u.rooms[len(u.rooms)-1]
			u.rooms = u.rooms[:len(u.rooms)-1]
		}
	}
}

// Notify sends a message to the Room, which then broadcasts it to all listeners
func (u *User) Notify(e Event) {
	for _, room := range u.rooms {
		if room.ID == e.TargetRoom.ID {
			room.OnNotify(e)
		}
	}
}

// OnNotify is for when a user recieves a new message from a Room they are in
func (u *User) OnNotify(e Event) {
	fmt.Printf("%s recieved a message from %s ----- \n", u.Name, e.Sender.Name)
}

// NewUser creates an instance of the user struct
func NewUser(name string) User {
	return User{
		messenger: nil,
		Name:      name,
	}
}
