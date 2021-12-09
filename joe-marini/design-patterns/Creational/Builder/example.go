package main

import (
	"fmt"
	"log"
)

func main() {
	// Create a NotificationBuilder and use it to set properties
	bldr := newNotificationBuilder()

	// Use the builder to set some properties
	bldr.SetTitle("New Notification")
	bldr.SetSubTitle("This is a subtitle")
	bldr.SetIcon("icon.png")
	bldr.SetImage("image.png")
	bldr.SetPriority(5)
	bldr.SetMessage("This is a basic Notification with some text in it")
	bldr.SetType("alert")

	// Use the Build function to create a finished object
	notification, err := bldr.Build()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%+v", notification)
}
