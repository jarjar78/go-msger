// messenger
package messenger

import (
	mn "mynet"
	"net"
	"ui"
)

type Messenger struct {
}

func (messenger *Messenger) Communicate(udpAddr *net.UDPAddr, ui *ui.UI) {

	canal := make(chan string)
	rcanal := make(chan string)
	listener := new(mn.Reciever)
	sender := new(mn.Sender)
	go listener.Read(udpAddr, rcanal)
	go sender.Run(udpAddr, canal)
	ui.ShowWindow(canal, rcanal)
}
