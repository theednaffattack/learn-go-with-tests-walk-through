// first declare it as a package
package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bon Juor, "
const spanish = "Spanish"
const french = "French"

type HelloParams struct {
	name string
	lang string
}

// name string, language string
func Hello(arg HelloParams) string {
	// If supplied string is empty,
	// return default message.
	if arg.name == "" {
		arg.name = "World"
	}
	return greetingPrefix(arg.lang) + arg.name
}

func greetingPrefix(language string)(prefix string){
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main(){
	fmt.Println(Hello(HelloParams{name: "Eddie", lang: "English"}))
}