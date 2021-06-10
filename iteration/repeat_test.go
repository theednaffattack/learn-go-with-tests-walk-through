package iteration

import "testing"

func TestRepeat(t *testing.T){
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B){
	for index := 0; index < b.N; index++ {
		Repeat("a")
	}
}