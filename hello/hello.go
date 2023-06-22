package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const spanishPrefix = "Hola, "
const englishHelloPrefix = "Hello, "
const frenchPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	}

	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("world", ""))
}