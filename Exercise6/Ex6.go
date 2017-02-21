//Process pairs
package main

import (
	_ "bytes"
	"encoding/binary"
	"fmt"
	_ "log"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

//var serverAddr *net.UDPAddr
//var localIP string
var port = ":30042"
var msg = "Hello from the other side"
var conn *net.UDPConn

func check_error(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

// All credz Anders.
// Brukes til å sammenligne om det er en meling fra deg selv.
// I såfall ignorer denne.
func get_local_IP() string {
	conn, err := net.DialTCP("tcp4", nil, &net.TCPAddr{IP: []byte{8, 8, 8, 8}, Port: 53})
	check_error(err)
	defer conn.Close()

	localIP := strings.Split(conn.LocalAddr().String(), ":")[0]
	return localIP
}

func Init() (*net.UDPAddr, string) {
	localIP := get_local_IP()

	// setting up UDP server for broadcasting
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1"+port)
	check_error(err)
	fmt.Println("Server adress: ", serverAddr)
	fmt.Println("Local adress: ", localIP)

	return serverAddr, localIP
}

func spawn_new_terminal() {
	command := exec.Command("cmd", "/C go run Ex6.go")
	_ = command.Run()
}

func main() {
	IP := get_local_IP()
	var master bool = false
	var counter uint64 = 0
	fmt.Println("Local IP is: ", IP)
	udpaddr, _ := net.ResolveUDPAddr("udp", IP+":30005")
	connection, err := net.ListenUDP("udp", udpaddr)
	if err != nil {
		fmt.Println("Housten, we have a problem!")
	}

	fmt.Println("backup running")
	UDPmsg := make([]byte, 8)

	for !(master) {
		connection.SetReadDeadline(time.Now().Add(time.Second * 2))

		n, _, err := connection.ReadFromUDP(UDPmsg)

		if err == nil {
			counter = binary.BigEndian.Uint64(UDPmsg[0:n])

		} else {
			master = true
		}

	}
	fmt.Println(counter)
	connection.Close()

	fmt.Println("Connection was closed")
	//spawnNewTerminal()

	connection, _ = net.DialUDP("udp", nil, udpaddr)
	go func() {
		for {

			fmt.Println(counter)
			counter++
			binary.BigEndian.PutUint64(UDPmsg, counter)
			fmt.Println("Sending message: ", UDPmsg)
			_, _ = connection.Write(UDPmsg)

			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			connection.SetReadDeadline(time.Now().Add(time.Second * 2))

			n, _, err := connection.ReadFromUDP(UDPmsg)

			if err == nil {
				counter = binary.BigEndian.Uint64(UDPmsg[0:n])
				fmt.Println("Recieved message?")

			} else {
				master = true
				fmt.Println("Error in reading")
			}
		}
	}()

	for {
	}

}
