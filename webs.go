package main

import "net/http"
import "fmt"
import "flag"

func main() {

	defaultPortPtr := flag.String("p", "", "Port Number")
	
	flag.Parse()
	
	portNum := "9090"
	
	// Port Number
	if *defaultPortPtr != "" {
		portNum = *defaultPortPtr
	} else {
		portNum = "9090"
	}

	fmt.Println()
	fmt.Println("HTTP Server On :" + portNum + " Started .")
	fmt.Println()
	fmt.Println("Ver 1.0.0")
	panic(http.ListenAndServe(":"+portNum, http.FileServer(http.Dir("./"))))
}

