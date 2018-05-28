package main

// Comcurrent Scan Thread and lightweight..
// Outline .
// Read targets
// make directory
// IMPLEMENT A CHECK FOR HOSTS AND THEN MAKE!!
// How do i pass the ping into a functional and pass
import (
	"bufio"
	//"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	file, err := os.Open("targets.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		PingCheck(scanner.Text())

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	FileRead()

}

// This just needs to regex from the below to then parse that to scan the port..
func FileRead() {
	file, err := os.Open("hosts.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileread := bufio.NewScanner(file)
	for fileread.Scan() {
		PortScan(fileread.Text())

	}
	if err := fileread.Err(); err != nil {
		log.Fatal(err)
	}

}

func PortScan(ip string) {
	fmt.Printf("%s\n", ip)
	for port := 1; port < 65535; port++ {
		PortScan_Con(ip, port)

	}
}

func PortScan_Con(ip string, port int) {
	OpenPorts, _ := os.OpenFile("open.txt", os.O_APPEND, 0666)
	writer := bufio.NewWriter(OpenPorts)
	hostIP := fmt.Sprintf("%s:%d", ip, port)

	_, err := net.DialTimeout("tcp", hostIP, 20*time.Millisecond)

	if err != nil {
		fmt.Printf("Port %d is closed\n", port)
		//fmt.Fprintf(writer, "Closed %d\n", port)
		writer.Flush()

	} else {
		fmt.Printf("Port %d is open\n", port)
		fmt.Fprintln(writer, "Open %d", port)
		writer.Flush()
	}

}

func PingCheck(a string) {
	fileHandle, _ := os.OpenFile("hosts.txt", os.O_APPEND, 0666)
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	resolve, err := net.LookupIP(a)
	if err != nil {
		fmt.Println("Unable to Resolve name..\n")
	} else {
		fmt.Printf("%s Resovles to %s\n", a, resolve)
	}
	fmt.Fprintln(writer, resolve)
	writer.Flush()

}
