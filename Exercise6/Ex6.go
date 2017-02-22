package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const COUNTER_FILE = "counterSave.txt"
const PORT_FILE = "port.txt"

func main() {

	// Setup:
	fmt.Println("This is backup")
	localIP := "127.0.0.1"
	localPort := read_port_from_file()
	sendToPort := write_port_to_file(localPort)
	localAddr := localIP + ":" + localPort
	sendToAddr := localIP + ":" + sendToPort

	fmt.Println("this is sendToAddr: ", sendToAddr)
	udpAddr, err := net.ResolveUDPAddr("udp4", localAddr)
	check_error(err)

	connection, err := net.ListenUDP("udp4", udpAddr)
	check_error(err)

	fmt.Println("Listening on port ", localPort)
	defer connection.Close()

	masterChan := make(chan bool)

	// Ser om det er noen som sender "I am the master"
	go detect_spam(connection, masterChan)

	<-masterChan
	fmt.Println("I am master")
	backup()
	counterChan := make(chan int)
	go read_from_file(counterChan)
	time.Sleep(time.Second)
	go spam(sendToAddr)
	counter_and_write_to_file(counterChan)
}

func counter_and_write_to_file(counterChan chan int) {
	counter := <-counterChan
	for {
		counter++

		// Overskriver filen counter lagres i
		dataFile, err := os.Create(COUNTER_FILE)
		check_error(err)

		// Converterer int -> string og lagrer i filen
		dataFile.WriteString(strconv.Itoa(counter))
		dataFile.Close()
		fmt.Println(counter)
		time.Sleep(time.Millisecond * 200)
	}
}

func read_from_file(counterChan chan int) {
	var counter int

	// Dersom det ikke eksister en fil, sett counter til 0
	if _, err := os.Stat(COUNTER_FILE); os.IsNotExist(err) {
		counter = 0
	} else {

		// Leser fra fil og converterer til int
		data, err := ioutil.ReadFile(COUNTER_FILE)
		check_error(err)
		counter, _ = strconv.Atoi(string(data))
		fmt.Println(counter)
	}
	counterChan <- counter
}

func backup() {
	// DENNE MÅ ENDRES!!!
	// Oppsett av terminal på mac
	// arg := "Tell application \"Terminal\"\n set newTab to do script [\"cd Desktop/\"]\n do script [\"go run ov6.go\"] in newTab\n end Tell"
	command := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run Ex6.go")
	err := command.Run()
	check_error(err)

	fmt.Println("backup should be spawned")
}

func spam(remoteAddr string) {
	// Setup
	udpRemote, _ := net.ResolveUDPAddr("udp", remoteAddr)

	// Åpner kobling
	connection, err := net.DialUDP("udp", nil, udpRemote)
	check_error(err)
	defer connection.Close()

	for {
		// Dender melding 2 ganger i sekundet
		_, err := connection.Write([]byte("I am the master"))
		check_error(err)
		time.Sleep(time.Millisecond * 500)
	}
}

func detect_spam(connection *net.UDPConn, masterChan chan bool) {
	// Her lagres meldingen
	buffer := make([]byte, 2048)

	// Evig while løkke helt til det ikke registreres en melding
	for {
		// Timer for hvor lenge man skal vente på melding
		t := time.Now()
		connection.SetReadDeadline(t.Add(3 * time.Second))
		_, _, err := connection.ReadFromUDP(buffer)

		// Dersom error blir backup ny master og går ut av løkka
		if err != nil {
			fmt.Println("UDP timeout: ", err)
			masterChan <- true
			break
		}
	}
}

func write_port_to_file(portNum string) string {
	// String -> int
	portNumToFile, err := strconv.Atoi(portNum)
	check_error(err)
	portNumToFile++

	// Overskriver den gamle filen
	dataFile, err := os.Create(PORT_FILE)
	check_error(err)

	// Og legger inn port
	dataFile.WriteString(strconv.Itoa(portNumToFile))
	dataFile.Close()

	return strconv.Itoa(portNumToFile)
}

func read_port_from_file() string {
	// Dersom det ikke eksister en fil, returnerer vi start porten
	if _, err := os.Stat(PORT_FILE); os.IsNotExist(err) {
		return "30042"
	}

	// Leser fra fil
	data, err := ioutil.ReadFile(PORT_FILE)
	check_error(err)

	return string(data)
}

func check_error(err error) {
	// Standard error check
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}
