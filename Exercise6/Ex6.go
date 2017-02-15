//Process pairs
package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
  "os/exec"
  _ "bytes"
  _ "log"
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

func Recive_msg_UDP(channel chan){
	serverAddr, err := net.ResolveUDPAddr("udp", port)
	check_error(err)

	conn, err := net.ListenUDP("udp", serverAddr)
	check_error(err)
	defer conn.Close()

	buffer := make([]byte, 1024)

	n, address, err := conn.ReadFromUDP(buffer)
	check_error(err)
	fmt.Println("Got message from ", address, " with n = ", n)
	if n > 0 {
		fmt.Println("From address: ", address, " got message: ", string(buffer[0:n]))
    channel <- true
	}
	fmt.Println("Listening...")
	time.Sleep(100 * time.Millisecond)
}

func Broadcast_UDP(serverAddr *net.UDPAddr) {
	localAddr, err := net.ResolveUDPAddr("udp", ":0")
	check_error(err)

	conn, err := net.DialUDP("udp", localAddr, serverAddr)
	check_error(err)
	defer conn.Close()

	fmt.Println("Sending message...")
	_, err = conn.Write([]byte(msg))
	check_error(err)
	time.Sleep(1000 * time.Millisecond)
}

func main() {
  primary := false
  timer := false

  channel := make(chan bool)

  timer := time.NewTimer(5*time.Second)

  for {
    if <-channel{
      timer := time.NewTimer(5*time.Second)
    }
    if <-timer{
      //you are primary
    }
  }

  serverAddr, _ := Init()


  cmd := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run Ex6.go")
  err := cmd.Start()
  check_error(err)

  IP := get_local_IP()
  fmt.Println(IP)


}
