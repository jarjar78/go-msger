// mynet
package mynet

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

type Sender struct {
	bytes []byte
	conn  net.Conn
	err   error
}
type Reciever struct {
	bytes, otherBytes []byte
	conn              *net.UDPConn
	err               error
}

func Login(name string, udpAddr *net.UDPAddr) {
	resp, err := http.Get("http://127.0.0.1:8889/userup?name=" + name + "&port=" + udpAddr.String())
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || string(body) != "ok" {
		fmt.Println(err)
	}
}

func (receiver *Reciever) Read(udpAddr *net.UDPAddr, rcanal chan string) {
	var from *net.UDPAddr
	var len int
	receiver.otherBytes = receiver.bytes
	receiver.conn, receiver.err = net.ListenUDP("udp", udpAddr)
	if receiver.err != nil {
		fmt.Println(receiver.err)
	}
	for {
		receiver.bytes = make([]byte, 2048)
		len, _, _, from, receiver.err = receiver.conn.ReadMsgUDP(receiver.bytes, receiver.otherBytes)
		fmt.Println(string(receiver.bytes[:len]), " from ", *from)
		rcanal <- string(receiver.bytes[:len])

		if receiver.err != nil {
			fmt.Println(receiver.err)
			continue
		}
	}

}

func (sender *Sender) Run(udpAddr *net.UDPAddr, c chan string) {
	for {
		msg := <-c
		sender.Send(udpAddr, msg)
	}
}

func (sender *Sender) Send(udpAddr *net.UDPAddr, msg string) {
	sender.conn, sender.err = net.Dial("udp", ":1200")
	if sender.err != nil {
		fmt.Println(sender.err)
	}
	sender.conn.Write([]byte(msg))
}

func GetUserlist(name string) []string {
	resp, err := http.Get("http://127.0.0.1:8889/list?name=" + name)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(body)
	list := []string{"vasya", "petya", "manya"}
	return list
}
