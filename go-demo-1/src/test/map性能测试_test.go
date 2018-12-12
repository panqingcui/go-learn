package test

import (
	"fmt"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	fmt.Println("ddd {]", b.N)
	for i := 0; i < b.N; i++ {
		test()
	}
}
func BenchmarkTestCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCap()
	}
}
func test() map[int]int {
	m := make(map[int]int)
	for i := 0; i < 100000; i++ {
		m[i] = i
	}
	return m
}
func testCap() map[int]int {
	m := make(map[int]int, 100000)
	for i := 0; i < 100000; i++ {
		m[i] = i
	}
	return m
}
