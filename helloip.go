package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	eth, err := net.InterfaceByName("eth0")
	if err != nil {
		fmt.Println(err)
	}
	ip, _ := eth.Addrs()
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "<h1>Hello World from %v!</h1>\n<b>My IP address is %v</b>\n", hostname, ip[0])
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Started, serving at 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
