package mcpiapi

import (
	"bufio"
	"fmt"
	//	"log"
	"net"
	"strings"
	"time"
)

// response is sent back over the reply channel.
type response struct {
	text string
	err  error
}

// message is sent over the request channel.
type message struct {
	text   string
	rspBuf chan response
	rcv    bool
}

// Connection provides access to the Minecraft API.
type Connection struct {
	conn   net.Conn
	msgBuf chan message
	name   string
}

// Chat returns the chat object.
func (obj *Connection) Chat() Chat {
	return Chat(obj.msgBuf)
}

// World returns the world object.
func (obj *Connection) World() World {
	return World(obj.msgBuf)
}

// Camera returns the camera object.
func (obj *Connection) Camera() Camera {
	return Camera(obj.msgBuf)
}

// Player returns the player object.
func (obj *Connection) Player() Player {
	p := Player{obj.msgBuf, obj.name}
	return Player(p)
}

// Events returns the events object.
func (obj *Connection) Events() Events {
	return Events(obj.msgBuf)
}

// Open establishes a connection to the given host over port 4711 with player name.
func (obj *Connection) Open(host string, name string) error {
	var err error
	obj.conn, err = net.Dial("tcp", host+":4711")
	obj.name = name
	if err != nil {
		return err
	}
	obj.msgBuf = make(chan message, 1024)
	go obj.process()
	return nil
}

// Close closes the connection to the host.
func (obj *Connection) Close() error {
	msg := message{"###", make(chan response, 1), false}
	defer close(msg.rspBuf)
	obj.msgBuf <- msg
	_ = <-msg.rspBuf
	return obj.conn.Close()
}

// send sends a message to the host over the open connection.
func (obj *Connection) send(msg string) error {
	_, err := fmt.Fprintf(obj.conn, "%s\n", msg)
	return err
}

// receive reads a message from the host over the open connection.
func (obj *Connection) receive() (string, error) {
	s, e := bufio.NewReader(obj.conn).ReadString('\n')
	s = strings.TrimSpace(s)
	return s, e
}

// process reads the request channel and processes commands.
func (obj *Connection) process() {
	for msg := range obj.msgBuf {
		if msg.text == "XXX" {
			// log.Print("Close request received")
			msg.rspBuf <- response{"OK", nil}
			close(obj.msgBuf)
			return
		}
		err := obj.send(msg.text)
		// log.Printf("snd [%s] err [%v]", msg.text, err)
		if msg.rcv {
			rsp, err := obj.receive()
			// log.Printf("rcv [%s] err [%v]", rsp, err)
			msg.rspBuf <- response{rsp, err}
		} else {
			msg.rspBuf <- response{"", err}
		}
		// give the players a little CPU time on the pi
		time.Sleep(time.Millisecond)
	}
}

// object is a type that allows us to provide the send and sendReceive methods
// to all of the Minecraft API objects - world, player, etc.
type object chan message

// send sends a command over the request channel and waits for the reply.
func (obj object) send(s string) error {
	rspBuf := make(chan response, 1)
	defer close(rspBuf)
	msg := message{s, rspBuf, false}
	obj <- msg
	//fmt.Println(msg)
	rsp := <-rspBuf
	return rsp.err
}

// sendReceive sends a command over the request channel and receives
// a reply containing data returned by the Minecraft server.
func (obj object) sendReceive(s string) (string, error) {
	rspBuf := make(chan response, 1)
	defer close(rspBuf)
	msg := message{s, rspBuf, true}
	obj <- msg
	//fmt.Println(msg)
	rsp := <-rspBuf
	return rsp.text, rsp.err
}
