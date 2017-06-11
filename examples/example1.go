package main

import (
	"fmt"
	"github.com/shomali11/commander"
)

func main() {
	properties, isMatch := commander.NewCommand("").Match("ping")
	fmt.Println(isMatch)
	fmt.Println(properties)

	properties, isMatch = commander.NewCommand("").Match("")
	fmt.Println(isMatch)
	fmt.Println(properties)

	properties, isMatch = commander.NewCommand("ping").Match("ping")
	fmt.Println(isMatch)
	fmt.Println(properties)

	properties, isMatch = commander.NewCommand("ping").Match("pong")
	fmt.Println(isMatch)
	fmt.Println(properties)

	properties, isMatch = commander.NewCommand("echo <word>").Match("echo")
	fmt.Println(isMatch)
	fmt.Println(properties)
}
