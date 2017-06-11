package main

import (
	"fmt"
	"github.com/shomali11/commander"
)

func main() {
	properties, isMatch := commander.NewCommand("echo <word>").Match("echo hey")
	fmt.Println(isMatch)
	fmt.Println(properties.StringParam("word", ""))

	properties, isMatch = commander.NewCommand("repeat <word> <number>").Match("repeat hey 5")
	fmt.Println(isMatch)
	fmt.Println(properties.StringParam("word", ""))
	fmt.Println(properties.IntegerParam("number", 0))

	properties, isMatch = commander.NewCommand("repeat <word> <number>").Match("repeat hey")
	fmt.Println(isMatch)
	fmt.Println(properties.StringParam("word", ""))
	fmt.Println(properties.IntegerParam("number", 0))
}
