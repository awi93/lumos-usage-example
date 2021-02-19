package handler

import (
	"fmt"
	"github.com/dimall-id/lumos/event"
)

type CommandEventHandler struct {}

func (c *CommandEventHandler) Handler (message event.ConsumerMessage) error {
	fmt.Println(message.Headers)
	fmt.Println(message.Topic,message.Value)
	return nil
}
