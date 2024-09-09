package main

import (
	"github.com/rs/zerolog/log"

	"github.com/madmuzz05/golang-assigment-2/api/server"
)


func main(){
	err := server.InitServer()
	if err != nil {
		log.Panic().Err(err)
	}
}