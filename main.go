package main

import "fmt"

const spanish = "Spanish"
const english = "English"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	prefix := englishHelloPrefix
	if lang == spanish {
		prefix = spanishHelloPrefix
	} else if lang == french {
		prefix = frenchHelloPrefix
	}
	if name == "" {
		name = "World"
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("World", english))
}
