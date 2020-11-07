package main

import (
	"./server"
)

func startServer() *server.Server {
	serv := server.Server{};
	serv.Start();
	return &serv
}

func main(){
	//serv:= server.NewServer("localhost", "80")
	startServer()
}
