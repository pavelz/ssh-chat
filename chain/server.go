package chain

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type ChainServer struct {
	port uint16
}

/*
	Accept a connection
	Pass to handler
*/

func (c *ChainServer) Serve() {
	ln, err := net.Listen("tcp", fmt.Sprintf("%d", c.port))
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()

		if err != nil {
			conn.Close()
			log.Print(err)
		}

		go c.Handle(conn)
	}
}

/*
	Listen
	Call this method form go routine, apparently go
*/

func (c *ChainServer) Handle(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		cmd, err := reader.ReadString('\n')
		if err != nil {
			conn.Close()
			log.Print(err)
			return
		}
		fmt.Println(cmd)
	}
}
