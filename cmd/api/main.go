package main

import (
	"crypto/rsa"
	"log"
	"os"
)

func generateKeys() []string {
	rsa.GenerateKey()
	return nil
}

func main() {
	log := log.New(os.Stdout, "AUTH : ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	log.Println("Starting...")
}
