package main

import(
	"os"
	"log"
	"fmt"
	"unicode"
)
func main() {
	text, err := os.ReadFile("input")
	if err != nil{
		log.Fatalf("unable to read file %v", err)
	}
}
