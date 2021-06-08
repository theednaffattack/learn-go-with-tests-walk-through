// first declare it as a package
package main

import "fmt"

const englishPrefix = "Hello, "

func Hello(name string) string {
	// If supplied string is empty,
	// return default message.
	if name == "" {
		name = "World"
	}

	return englishPrefix + "World"

}

func main(){
	fmt.Println(Hello("Eddie"))
}