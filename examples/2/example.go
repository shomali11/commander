package main

import (
	"fmt"
	"github.com/shomali11/commander"
)

func main() {
	tokens := commander.NewCommand("echo {word} <sentence>").Tokenize()
	for _, token := range tokens {
		fmt.Println(token)
	}
}
