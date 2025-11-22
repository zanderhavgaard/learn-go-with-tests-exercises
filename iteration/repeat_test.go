package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"
	if repeated != expected {
		t.Errorf("expected %q got %q", expected, repeated)
	}
}

// run bechamrks with 'go test -bench=.'
// 'go test -benchmem -bench=.' includes memory benchmarking info
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeated := Repeat("foo", 3)
	fmt.Println(repeated)
	// output: foofoofoo
}
