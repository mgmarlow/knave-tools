package main

import (
	"math/rand"
)

func Roll(sides int) int {
	return rand.Intn(sides) + 1
}

func Includes[T comparable](arr []T, want T) bool {
	for _, v := range arr {
		if v == want {
			return true
		}
	}

	return false
}

func Sample[T any](arr []T) T {
	index := rand.Intn(len(arr))
	return arr[index]
}

func Map[T any, R any](collection []T, iteratee func(T, int) R) []R {
	result := make([]R, len(collection))

	for i, item := range collection {
		result[i] = iteratee(item, i)
	}

	return result
}

func Min(arr []int) int {
	cur := arr[0]

	for _, v := range arr {
		if cur > v {
			cur = v
		}
	}

	return cur
}
