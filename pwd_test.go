package main

import (
	"os"
	"testing"
)

func TestPwd(t *testing.T) {
	expected, _ := os.Getwd()
	if observed := pwd(); observed != expected {
		t.Fatalf("pwd() = %v, want %v", observed, expected)
	}
	t.Logf("PASS: pwd")
}

func BenchmarkPwd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pwd()
	}
}
