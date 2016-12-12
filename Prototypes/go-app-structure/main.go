package main

import (
	"fmt"

	"github.com/dotbugfix/GoNFS4/Prototypes/go-app-structure/config"
	"github.com/dotbugfix/GoNFS4/Prototypes/go-app-structure/ui"
	"path/filepath"
	"os"
	"log"
)

func main() {
	log.Println("This is a test")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	config.PrintConfig()
	ui.StartUIServer()
}
