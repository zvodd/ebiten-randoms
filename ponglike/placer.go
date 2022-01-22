package main

import (
	"math/rand"
)

func Shuffled(size int) []int {
	flatlist := make([]int, size)

	for i := 0; i < size; i++ {
		flatlist[i] = i
	}
	for i := size; i > 0; i-- {
		pick := rand.Intn(i)
		C := flatlist[i-1]
		P := flatlist[pick]
		flatlist[i-1] = P
		flatlist[pick] = C
	}
	return flatlist
}
