package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"time"
)

//Server to Client
const (
	HeartBeatResp    byte = 0x0C
	NewConnection    byte = 0x0A
	NewDisconnection byte = 0x0B
)

//Server
type Server struct {
	mConn    net.Listener
	mClients []*Client
}

//Init Server
func InitServer(port int) *Server {
	addr := string(":" + strconv.Itoa(port))
	server, err := net.Listen("tcp", addr)

	if err != nil {
		fmt.Printf("Cannot start socket.\n")
		return nil
	}

	fmt.Printf("Socket Running on port %s\n", addr)

	ins := &Server{mConn: server}

	return ins
}

func handleClient(conn net.Conn, server *Server, queue *Queue) {
	fmt.Printf("Connection from (%s)\n", conn.RemoteAddr().String())
	reader := bufio.NewReader(conn)
	var name string
	var client *Client

	read, _ := reader.ReadBytes('\n')
	if len(read) == 0 {
		goto dis
	}

	if read[0] == Log {
		name = string(read[1 : len(read)-1])
		fmt.Printf("The name is %s\n", name)

		client = &Client{mConn: conn, mName: name}
		aqwarg := NewQwarg(addClient, server, client)
		AddtoQueue(queue, aqwarg)

		cmsg := fmt.Sprintf("%c%s\000", NewConnection, client.mName)
		for _, c := range server.mClients {
			if c != client {
				c.mConn.Write([]byte(cmsg))
			}
		}

		//loop
		for {
			read, _ := reader.ReadBytes('\n')

			if len(read) == 0 {
				goto dis
			}

			switch read[0] {
			case HeartBeat:
				fmt.Printf("%s heartbeating\n", name)
				//client.mConn.Write([]byte{HeartBeatResp, 0x00})
				break

			case Message:
				msg := read[1 : len(read)-1]
				snd := fmt.Sprintf("%c%s\000%s\000", Message, client.mName, msg)
				fmt.Printf("%s\n", msg)

				for _, c := range server.mClients {
					if c != client {
						c.mConn.Write([]byte(snd))
					}
				}
				break

			case Disconnection:
				goto dis

			default:
				goto dis
			}
			time.Sleep(time.Nanosecond * 5)
		}
	}

dis:
	if client == nil {
		conn.Close()
		fmt.Printf("(%s) disconnected\n", conn.RemoteAddr().String())
		return
	}

	rqwarg := NewQwarg(removeClient, server, client)
	AddtoQueue(queue, rqwarg)
	conn.Close()
	fmt.Printf("(%s) disconnected\n", conn.RemoteAddr().String())

	dmsg := fmt.Sprintf("%c%s\000", NewDisconnection, client.mName)
	for _, c := range server.mClients {
		if c != client {
			c.mConn.Write([]byte(dmsg))
		}
	}
}

//Server Mainloop
func ServerMainloop(server *Server, queue *Queue) {
	for {
		cli, err := server.mConn.Accept()
		if err != nil {
			fmt.Printf("Cannot connect client\n")
		} else {
			go handleClient(cli, server, queue)
		}
	}
}
