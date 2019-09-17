package gochat

import (
	"fmt"
)

type (
	// Event contains a record of a message event, with a sender and a message
	Event struct {
		Sender     User
		TargetRoom Room
		Message    Message
	}

	// Observer defines a listener in a Room
	Observer interface {
		OnNotify(Event)
	}

	// Notifier defines a speaker in a Room
	Notifier interface {
		Register(Observer)
		DeRegister(Observer)
		Notify(Event)
	}
)

// UserListIterator is used to iterate over the entire chat log effieciently
type UserListIterator struct {
	room Room
	curr int
	err  error
}

// HasNext checks to see if there are more elements left in the iterator
func (i *UserListIterator) HasNext() bool {
	return cap(i.room.users) > i.curr
}

// Get returns the current chat log and increments the iterator position
func (i *UserListIterator) Get() User {
	i.curr++
	return i.room.users[i.curr-1]
}

// NewUserListIterator takes a Room to generate an iterator from
func NewUserListIterator(room Room) UserListIterator {
	return UserListIterator{
		room: room,
		curr: 0,
		err:  nil,
	}
}

// Room is a struct to hold all relevant information that people may need to communicate with each other
type Room struct {
	ID    string
	users []User
	log   []Event
}

// Register adds a user to the list of listeners in a Room
func (r *Room) Register(u User) {
	r.users = append(r.users, u)
	// Register Room to allow user to send messages to the Room
	u.Register(*r)
}

// Deregister removes a user from the list of listeners in a Room
func (r *Room) Deregister(name string) {
	for i, user := range r.users {
		if user.Name == name {
			r.users[i] = r.users[len(r.users)-1]
			r.users = r.users[:len(r.users)-1]
			// Remove Room from the user's list of Rooms
			user.Deregister(*r)
		}
	}
}

// Notify sends out a message event to all listening users
func (r *Room) Notify(e Event) {
	for _, user := range r.users {
		if user.Name != e.Sender.Name {
			user.OnNotify(e)
		}
	}
}

// OnNotify has the Room recieve a message from one user, then broadcast it to all others
func (r *Room) OnNotify(e Event) {
	for _, user := range r.users {
		if user.Name != e.Sender.Name {
			user.OnNotify(e)
		}
	}
}

// PrintAllRoomUsers uses the iterator pattern to iterate over a Room's whole user listing
func (r *Room) PrintAllRoomUsers() {
	iter := NewUserListIterator(*r)

	for iter.HasNext() {
		u := iter.Get()
		fmt.Printf("%s is currently in room %s\n", u.Name, r.ID)
	}
}

// NewRoom creates a new room using a passed in ID value
func NewRoom(id string) Room {
	return Room{
		ID:    id,
		users: make([]User, 0),
		log:   make([]Event, 0),
	}
}
