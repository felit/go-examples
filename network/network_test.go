package network

import (
	"net"
	"testing"
)

func TestDial(t *testing.T){
	conn,err := net.Dial("tcp", "localhost:8000")
	//err有哪些
	if(err!= nil) {
		println(err)
	}
	conn.Close()
}
func TestListening(t *testing.T) {
	ln,err :=net.Listen("tcp", ":8080")
	if err!= nil {

	}
	for {
		conn ,err :=ln.Accept()
		if err != nil {
			continue
		}
		conn.Close()
	}
}
