package main

import (
	"fmt"

	chat "github.com/retwolf/gochat"
)

func main() {
	dogRoom := chat.NewRoom("dog")

	tom := chat.NewUser("Tom")
	bob := chat.NewUser("Bob")

	fmt.Println("Registering users to Room using Observer pattern")
	fmt.Println()

	dogRoom.Register(tom)
	dogRoom.Register(bob)

	tomText := chat.NewTextMessage("Testing123")

	fmt.Println("Sending message to Room using combination of strategy and observer pattern")
	fmt.Println()
	tom.SendMessage(&tomText, dogRoom)

	bobSticker := chat.NewStickerMessage()
	bob.SendMessage(&bobSticker, dogRoom)

	fmt.Println("Listing all Room users using Iterator pattern")
	fmt.Println()
	dogRoom.PrintAllRoomUsers()
}
