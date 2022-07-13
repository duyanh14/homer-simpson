package main

import (
	"golang-course/server"
)

func main() {
	//start new server
	s, err := server.NewServer()
	if err != nil {
		panic(err)
	}
	s.Init()

	if err := s.ListenHTTP(); err != nil {
		panic(err)
	}
}
