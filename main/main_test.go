package main

import (
	"fmt"
	"testing"
)

func TestRandomStuff(t *testing.T) {
	a := make([]int, 2, 5)

	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}
}
