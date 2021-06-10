package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T){
	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

var testConstant = "b"

func ExampleRepeat(){
	repeated := Repeat(testConstant,6)
	fmt.Println(repeated)
	// Output: bbbbbb
}

func BenchmarkRepeat(b *testing.B){
	for index := 0; index < b.N; index++ {
		Repeat("a", 8)
	}
}