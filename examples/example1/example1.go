package main

import (
	"fmt"

	"github.com/romnnn/flatbson"
)

func run() string {
	return flatbson.Shout("This is an example")
}

func main() {
	fmt.Println(run())
}
