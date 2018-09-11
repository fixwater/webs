package main

import "net/http"
import "fmt"

func main() {
	fmt.Println()
	fmt.Println("HTTP Server On :9090 Started .")
	fmt.Println()
	fmt.Println("Ver 1.0.0")
	panic(http.ListenAndServe(":9090", http.FileServer(http.Dir("./"))))
}

