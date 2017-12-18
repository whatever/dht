package main

import (
	"fmt"
	"net"
	"time"
)

// ...
func ReadCommand(conn *net.UDPConn) string {
	buffer := make([]byte, 1024)
	size, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		return "<This sucks>"
	}
	return string(buffer[:size])
}

// Return a channel that listens to a
func ListenerChannel(conn *net.UDPConn) chan string {
	buffer := make([]byte, 1024)
	c := make(chan string)

	go func() {
		for {
			size, _, err := conn.ReadFromUDP(buffer)

			if err != nil {
				break
			}

			c <- string(buffer[:size])
		}
	}()

	return c
}

//
func main() {
	// Server address
	address, _ := net.ResolveUDPAddr("udp", ":9090")
	connection, _ := net.ListenUDP("udp", address)
	defer connection.Close()

	c := ListenerChannel(connection)

	for {
		fmt.Println(<-c)
	}

	fmt.Println("x_x")
	time.Sleep(1)
}
