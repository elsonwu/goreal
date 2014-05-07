package goio

import (
	"sync"
	"time"
)

type Client struct {
	Event
	Id            string
	UserId        string
	Messages      []*Message
	LastHandshake int64
	LifeCycle     int64
	lock          sync.RWMutex
}

func (self *Client) CleanMessages() {
	self.lock.Lock()
	defer self.lock.Unlock()

	//make a default cap value, it's enough for most use
	self.Messages = make([]*Message, 0, 20)
}

func (self *Client) Receive(message *Message) {
	self.lock.Lock()
	defer self.lock.Unlock()

	// if there are too many messages, we limit the max to 80
	if len(self.Messages) >= 80 {
		self.Messages = self.Messages[1:]
	}

	self.Messages = append(self.Messages, message)
}

func (self *Client) Destroy() {
	self.Emit("destroy", nil)
}

func (self *Client) Handshake() {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.LastHandshake = time.Now().Unix()
}

func (self *Client) LifeRemain() int64 {
	remain := self.LifeCycle - (time.Now().Unix() - self.LastHandshake)
	if 0 >= remain {
		return 0
	}

	return remain
}

func (self *Client) IsLive() bool {
	return 0 < self.LifeRemain()
}
