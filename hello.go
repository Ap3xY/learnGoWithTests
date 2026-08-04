package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	switch language {
	case "Spanish":
		return spanishHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}

}

func main() {
	fmt.Println(Hello("Alex", "English"))
}
