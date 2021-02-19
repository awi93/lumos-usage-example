package event

import (
	"dimall.id/standard-template/event/handler"
	"github.com/dimall-id/lumos/event"
	"log"
)

func SetupHandler() {
	c := handler.CommandEventHandler{}
	err := event.AddCallback("COMMAND_EVENT", c.Handler)
	if err != nil {
		log.Fatal(err)
	}
}