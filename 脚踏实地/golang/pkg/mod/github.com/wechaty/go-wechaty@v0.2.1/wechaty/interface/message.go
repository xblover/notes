package _interface

import (
	file_box "github.com/wechaty/go-wechaty/wechaty-puppet/file-box"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"time"
)

type IMessageFactory interface {
	Load(id string) IMessage
	// Find find message in cache
	Find(query interface{}) IMessage
	// FindAll Find message in cache
	FindAll(query *schemas.MessageQueryFilter) []IMessage
}

type IMessage interface {
	Ready() (err error)
	IsReady() bool
	String() string
	Room() IRoom
	// Type get the type from the message.
	Type() schemas.MessageType
	From() IContact
	Text() string
	// Self check if a message is sent by self
	Self() bool
	Age() time.Duration
	// Message sent date
	Date() time.Time
	To() IContact
	// Get the recalled message
	ToRecalled() (IMessage, error)
	// Say reply a Text or Media File message to the sender.
	Say(textOrContactOrFileOrUrlOrMini interface{}) (IMessage, error)
	// Recall recall a message
	Recall() (bool, error)
	// MentionList get message mentioned contactList.
	MentionList() []IContact
	MentionText() string
	MentionSelf() bool
	Forward(contactOrRoomId string) error
	// ToFileBox extract the Media File from the Message, and put it into the FileBox.
	ToFileBox() (*file_box.FileBox, error)
	// ToImage extract the Image File from the Message, so that we can use different image sizes.
	ToImage() (IImage, error)
	// ToContact Get Share Card of the Message
	// Extract the Contact Card from the Message, and encapsulate it into Contact class
	ToContact() (IContact, error)
}
