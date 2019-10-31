package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

func main() {
	var address string

	network := "tcp"
	host := "google.com"
	port := "80"
	address = net.JoinHostPort(host, port)
	conn, err := net.Dial(network, address)

	if err != nil {
		fmt.Printf("Failed to connect.")
	} else {
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
		status, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Printf("Failure")
			return
		}

		fmt.Printf("Success:\n" + status + "\n")

		fmt.Printf("Making request to API mockup...\n")
		MakeRequest()
	}
}

func MakeRequest() {

	resp, err := http.Get("http://www.mocky.io/v2/5db9aec030000074cc5ee55f")

	if err != nil {
		fmt.Printf("Failed to make API call.\n")
	} else {
		if err == nil {
			var result map[string]string
			json.NewDecoder(resp.Body).Decode(&result)
			ip := result["ip"]

			fmt.Println(ip)
		}
	}
}
