package main

import "testing"

func BenchmarkReadManual(b *testing.B) {
	_ = readManual()
}

func BenchmarkReadReflect(b *testing.B) {
	_ = readReflect()
}
