package mcpiapi

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type response struct {
	text string
	err  error
}

type message struct {
	text   string
	rspBuf chan response
	rcv    bool
}

type Connection struct {
	conn   net.Conn
	msgBuf chan message
}

func (obj *Connection) Chat() Chat {
	return Chat(obj.msgBuf)
}

func (obj *Connection) World() World {
	return World(obj.msgBuf)
}

func (obj *Connection) Camera() Camera {
	return Camera(obj.msgBuf)
}

func (obj *Connection) Player() Player {
	return Player(obj.msgBuf)
}

func (obj *Connection) Events() Events {
	return Events(obj.msgBuf)
}

func (obj *Connection) Open(host string) error {
	var err error
	obj.conn, err = net.Dial("tcp", host+":4711")
	if err != nil {
		return err
	}
	obj.msgBuf = make(chan message, 1024)
	go obj.process()
	return nil
}

func (obj *Connection) Close() error {
	msg := message{"###", make(chan response, 1), false}
	defer close(msg.rspBuf)
	obj.msgBuf <- msg
	_ = <-msg.rspBuf
	return obj.conn.Close()
}

func (obj *Connection) send(msg string) error {
	_, err := fmt.Fprintf(obj.conn, "%s\n", msg)
	return err
}

func (obj *Connection) receive() (string, error) {
	s, e := bufio.NewReader(obj.conn).ReadString('\n')
	s = strings.TrimSpace(s)
	return s, e
}

func (obj *Connection) process() {
	for msg := range obj.msgBuf {
		if msg.text == "XXX" {
			log.Print("Close request received")
			msg.rspBuf <- response{"OK", nil}
			close(obj.msgBuf)
			return
		}
		err := obj.send(msg.text)
		log.Printf("snd [%s] err [%v]", msg.text, err)
		if msg.rcv {
			rsp, err := obj.receive()
			log.Printf("rcv [%s] err [%v]", rsp, err)
			msg.rspBuf <- response{rsp, err}
		} else {
			msg.rspBuf <- response{"", err}
		}
	}
}

type object chan message

func (obj object) send(s string) error {
	rspBuf := make(chan response, 1)
	defer close(rspBuf)
	msg := message{s, rspBuf, false}
	obj <- msg
	rsp := <-rspBuf
	return rsp.err
}

func (obj object) sendReceive(s string) (string, error) {
	rspBuf := make(chan response, 1)
	defer close(rspBuf)
	msg := message{s, rspBuf, true}
	obj <- msg
	rsp := <-rspBuf
	return rsp.text, rsp.err
}
