// first declare it as a package
package main

import "fmt"

func Hello() string {
	return "Hello world"
}

func main(){
	fmt.Println(Hello())
}