package main

import "fmt"

type ClientList map[string]bool

func main() {
	clientList := make(ClientList)
	fmt.Println(clientList)
	clientList["Marlon"] = true
	fmt.Println(clientList)
}
