package main

import (
	"github.com/EdgardoArriagada/hexagonal-go/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
