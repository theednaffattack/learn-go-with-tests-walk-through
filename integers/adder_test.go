package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T){
	sum := Add(2,2)
	expected := 4

	fmt.Println("ADDER TEST RUNNING")

	if sum != expected {
		t.Errorf("Expected '%d' but got '%d'", sum, expected)
	}
}

func ExampleAdd(){
	sum := Add(2,4)
	fmt.Println(sum)
	// Output: 6
}