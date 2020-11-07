package main

import (
	"fmt"
	"time"
)

//TO CHANGE THE PORT, CHANGE THIS CONSTANT
const PORT int = 3200

func dequeueLoop(queue *Queue) {
	for {
		Dequeue(queue)
	}

}

//DEBUG FUNCS
func printClientsConnected(server *Server) {
	for {
		fmt.Printf("%d clients connected\n", len(server.mClients))
		time.Sleep(time.Second * 2)
	}
}

//func(server *Server){fmt.Printf("%d clients connected", len(server.mClients)}

func main() {
	queue := InitQueue()
	server := InitServer(PORT)

	go dequeueLoop(queue)
	//go printClientsConnected(server)
	ServerMainloop(server, queue)

	return
}
