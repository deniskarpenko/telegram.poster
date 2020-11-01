package main

import (
	"fmt"
	"./server"
)

func main(){
	serv:= server.NewServer("localhost", 80)
	fmt.Println(serv)
}
