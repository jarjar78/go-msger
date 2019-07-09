// test
package main

import (
	"fmt"
	"math/rand"
	msg "messenger"

	"net"
	"strconv"
	"time"
	"ui"
)

func main() {
	var service string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	port := r.Intn((2000 - 1200)) + 1200
	service = ":" + strconv.Itoa(port)
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		fmt.Println(err)
	}

	messenger := new(msg.Messenger)
	ui := new(ui.UI)
	ui.UdpAddr = udpAddr
	ui.CreateApp()
	messenger.Communicate(udpAddr, ui)
	ui.App.Run()
}
