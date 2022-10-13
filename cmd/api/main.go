package main

import (
	"context"
	"log"
	"os"
)

type config struct {
}

func main() {
	log := log.New(os.Stdout, "AUTH : ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)

	if err := run(log); err != nil {
		log.Fatalf("error : %v", err)
	}

}

func run(log *log.Logger) error {
	log.Println("Starting...")
	ctx := context.Background()

}
