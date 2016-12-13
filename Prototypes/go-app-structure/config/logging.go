package config

import (
	"log"
)

// TestLogging Test various logging libraries
func TestLogging() {
	testGoNativeLogging()
}

func testGoNativeLogging() {
	// First set some flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("[go-app-structure] ")

	log.Println("This is a native log")
	// This causes an exit(1)
	// log.Fatal("Something fatal happened")

	// This causes a panic()
	// log.Panic("I'm going to panic now!")
}
