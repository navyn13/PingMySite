package main

import (
	"fmt"
	"log"

	"github.com/navyn13/PingMySite/cmd/api"
	"github.com/navyn13/PingMySite/configs"
)

func main() {

	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port))
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
