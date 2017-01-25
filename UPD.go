package main

import (
	"fmt"
	"net"
	"strings"
)


var IP string
var Port = 30000

func Check_Error(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}



// finding our local IP: 129.241.187.155
func Get_Local_IP() (string) {
	IP = ""
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: []byte{8, 8, 8, 8}, Port:53})
	go Check_Error(err)
	defer conn.Close()
	IP = strings.Split(conn.LocalAddr().String(), ":")[0]
	fmt.Println("Your local IP address is: ", IP)
	return IP
}


func main() {
	IP, err := net.ResolveUDPAddr("udp", net.JoinHostPort(Get_Local_IP(), "30000"))
	Check_Error(err)
	fmt.Println(IP)

}
