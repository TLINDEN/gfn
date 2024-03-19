package main

import "log"

func Exists[K comparable, V any](m map[K]V, v K) bool {
	if _, ok := m[v]; ok {
		return true
	}
	return false
}

func Die(err error) int {
	log.Fatal("Error: ", err.Error())

	return 1
}

// find an item in a list, generic variant
func Contains[E comparable](s []E, v E) bool {
	for _, vs := range s {
		if v == vs {
			return true
		}
	}

	return false
}

// Transpose a matrix, x=>y, y=>x
// via https://gist.github.com/tanaikech/5cb41424ff8be0fdf19e78d375b6adb8
func Transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)

	result := make([][]string, xl)

	for i := range result {
		result[i] = make([]string, yl)
	}

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}

	return result
}
