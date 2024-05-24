package main

import (
	"fmt"
	"github.com/AyakuraYuki/morse"
)

func main() {
	m := morse.New()

	message := "hello, world"
	morseCode, err := m.Encode(message)
	if err != nil {
		panic(err)
	}
	// "hello, world" is: ".... . .-.. .-.. --- --..-- / .-- --- .-. .-.. -.."
	fmt.Printf("%q is: %q\n", message, morseCode)

	encoded := "-. .- -- ."
	decoded, err := m.Decode(encoded)
	if err != nil {
		panic(err)
	}
	// "-. .- -- ." is: "NAME"
	fmt.Printf("%q is: %q\n", encoded, decoded)
}
