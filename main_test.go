package main

import (
	"testing"
)

func BenchmarkProgramCustom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Program("./customPath")
	}
}

func BenchmarkProgramNull(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Program("")
	}
}