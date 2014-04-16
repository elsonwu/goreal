package goreal

import (
	"time"
)

type Client struct {
	Event
	Id            string
	User          *User
	Msg           chan *Message
	LastHandshake int64
}

func (self *Client) Receive(message *Message) {
	self.Msg <- message
}

func (self *Client) Destory() {
	defer close(self.Msg)

	self.Emit("destory", &Message{
		EventName: "destory",
		Data:      self.Id,
	})
}

func (self *Client) UpdateActiveTime() {
	self.LastHandshake = time.Now().Unix()
}
