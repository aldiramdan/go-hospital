package main

import (
	"log"
	"os"

	"github.com/aldiramdan/hospital/configs"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if err := configs.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
