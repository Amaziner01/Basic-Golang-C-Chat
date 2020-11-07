package main

import "net"

//Client to Server
const (
	Log           byte = 0x01
	HeartBeat     byte = 0x02
	Message       byte = 0x03
	Disconnection byte = 0x04
)

//Client
type Client struct {
	mConn net.Conn
	mName string
}

func addClient(server *Server, client *Client) {
	server.mClients = append(server.mClients, client)
}

func removeClient(server *Server, client *Client) {
	a := &server.mClients
	if len(*a) > 0 {
		index := 0
		for i, o := range *a {
			if o == client {
				index = i
				break
			}
		}

		(*a)[index] = (*a)[len(*a)-1]
		(*a)[len(*a)-1] = nil
		(*a) = (*a)[0 : len(*a)-1]
	}
}
