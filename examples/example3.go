package main

import (
	"fmt"
	"github.com/shomali11/commander"
)

func main() {
	fmt.Println(commander.IsParameter("<value>"))
	fmt.Println(commander.IsParameter("<123>"))
	fmt.Println(commander.IsParameter("<value123>"))
	fmt.Println(commander.IsParameter("value>"))
	fmt.Println(commander.IsParameter("<value"))
	fmt.Println(commander.IsParameter("value"))
	fmt.Println(commander.IsParameter(""))
}
