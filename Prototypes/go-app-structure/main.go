package main

import (
	"fmt"

	"github.com/dotbugfix/GoNFS4/Prototypes/go-app-structure/config"
	"github.com/dotbugfix/GoNFS4/Prototypes/go-app-structure/ui"
)

const cs string = "constant"

func for_loop() {
	i := 10
	for i > 0 {
		fmt.Println(i)
		i--
	}

	for j := 1; j < 10; j++ {
		fmt.Println(j)
	}

}

func main() {
	const num int = 10

	for_loop()
	fmt.Println(cs, num)
	var s string = "initial"
	fmt.Println(s)

	var a, b int = 1, 2
	fmt.Println(a, b)

	f := "short"
	fmt.Println(f)

	fmt.Println("Hello" + "World!")
	fmt.Println("1+1=", 1+1)
	fmt.Println(true && true)
	fmt.Println(!true)

	config.PrintConfig()
	ui.StartUIServer()

}
