package main

import (
	"log"

	vfileutil "github.com/waterdudu/villa/fileutil"
)

func main() {

	lines, err := vfileutil.ReadLines("a.go")
	if err != nil {
		log.Println("Reade a.go error", err)
	}
	log.Println(lines)
}
