package main

import (
	"fmt"
	"github.com/shomali11/commander"
)

func main() {
	properties, isMatch := commander.NewCommand("ping").Match("ping")
	fmt.Println(isMatch)
	fmt.Println(properties)

	properties, isMatch = commander.NewCommand("ping").Match("pong")
	fmt.Println(isMatch)
	fmt.Println(properties)

	properties, isMatch = commander.NewCommand("echo {word}").Match("echo hello world!")
	fmt.Println(isMatch)
	fmt.Println(properties.StringParam("word", ""))

	properties, isMatch = commander.NewCommand("echo <sentence>").Match("echo hello world!")
	fmt.Println(isMatch)
	fmt.Println(properties.StringParam("sentence", ""))

	properties, isMatch = commander.NewCommand("repeat {word} {number}").Match("repeat hey 5")
	fmt.Println(isMatch)
	fmt.Println(properties.StringParam("word", ""))
	fmt.Println(properties.IntegerParam("number", 0))

	properties, isMatch = commander.NewCommand("repeat {word} {number}").Match("repeat hey")
	fmt.Println(isMatch)
	fmt.Println(properties.StringParam("word", ""))
	fmt.Println(properties.IntegerParam("number", 0))

	properties, isMatch = commander.NewCommand("search <stuff> {size}").Match("search hello there everyone 10")
	fmt.Println(isMatch)
	fmt.Println(properties.StringParam("stuff", ""))
	fmt.Println(properties.IntegerParam("size", 0))
}
