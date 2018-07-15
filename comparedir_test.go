package main

import "testing"

func TestComparedir(t *testing.T) {
	expected := "pwd"
	if _, observed := comparedir("../"); observed != expected {
		t.Fatalf("comparedir() = %v, want %v", observed, expected)
	}
	t.Logf("PASS: pwd")
}

func BenchmarkComparedir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		comparedir("../")
	}
}
