package main

import (
	"fmt"

	"log"
	"os"
	"path/filepath"

	"github.com/dotbugfix/GoNFS4/Prototypes/go-app-structure/config"
	_ "github.com/dotbugfix/GoNFS4/Prototypes/go-app-structure/ui"
)

func main() {
	log.Println("This is a test")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	config.PrintConfig()
	config.TestLogging()
	// ui.StartUIServer()
}
