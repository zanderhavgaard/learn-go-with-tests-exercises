// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

package main

import "fmt"

const (
	englishHello  string = "Hello"
	japaneseHello string = "こんにちは"
	italianHello  string = "Boungiorno"
	japanese      string = "にほんじん"
	italian       string = "italiano"
)

func Hello(name string, language string) string {
	// default to "World" if no name is provided
	if name == "" {
		name = "World"
	}

	var helloString string

	switch language {
	case japanese:
		helloString = name + ", " + japaneseHello + "!"
	case italian:
		helloString = italianHello + ", " + name + "!"
	default:
		helloString = englishHello + ", " + name + "!"
	}

	return helloString
}

func main() {
	fmt.Println(Hello("world", ""))
}
