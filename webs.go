package main

import (
	"net/http"
	"os"
)
import "fmt"
import "flag"


func main() {

	PortPtr    := flag.String("p", "", "Port Number")
	vFlagPtr   := flag.Bool("v", false, "Print Version")
	dirFlagPtr := flag.String("d", "./", "Set Serve Dir")

	flag.Parse()

	vValue   := *vFlagPtr
	dirValue := *dirFlagPtr
	verStr   := "1.0.0"

	//Print Version
	if vValue  {
		fmt.Println("Webs Server : V" + verStr)
		os.Exit(0)
	}


	
	portNum := "9090"
	
	// Port Number
	if *PortPtr != "" {
		portNum = *PortPtr
	} else {
		portNum = "9090"
	}

	fmt.Println()
	fmt.Println("HTTP Server On :" + portNum + " Started .")
	fmt.Println()
	fmt.Println("Dir is " + dirValue);
	fmt.Println()
	fmt.Println("Ver " + verStr)

	panic(http.ListenAndServe(":"+portNum, http.FileServer(http.Dir(dirValue))))

}

