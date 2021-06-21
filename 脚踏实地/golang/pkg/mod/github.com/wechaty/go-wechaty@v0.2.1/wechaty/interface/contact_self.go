package _interface

import file_box "github.com/wechaty/go-wechaty/wechaty-puppet/file-box"

type IContactSelfFactory interface {
	IContactFactory
}

type IContactSelf interface {
	IContact
	SetAvatar(box *file_box.FileBox) error
	// QRCode get bot qrcode
	QRCode() (string, error)
	SetName(name string) error
	// Signature change bot signature
	Signature(signature string) error
}
