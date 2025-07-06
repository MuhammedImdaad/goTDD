package basics

import "testing"

func TestIteration(t *testing.T) {
	sum := Repeat("A", 5)
	expected := "AAAAA"

	if sum != expected {
		t.Errorf("expected '%q' but got '%q'", expected, sum)
	}
}

func BenchmarkRepeat(b *testing.B) { // go test -bench=.
	//... setup ...
	for b.Loop() { // Loop() returns true as long as the benchmark should continue running.
		//... code to measure ...
		Repeat("a", 10)
	}
	//... cleanup ...
}
