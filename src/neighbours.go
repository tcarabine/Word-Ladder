package main

import (
	"strings"
)

func neighbours (dict *map[string]bool, word string) []string {
	var ret = []string
	wordSlice = strings.Split(word, "")

	for compare := range dictionary {
		if word != compare {

			compareSlice = strings.Split(compare, "")
			diff := 0

			for i := 0; i < length; i++ {
				if wordSlice[i] != compareSlice[i] {
					diff++
				}
			}
			if (diff == 1) && !(dict[compare]) {
				// If there is only one letter different and the value in the dictionary is false - not visited, add to return slice
				ret = append(ret,compare)
			}
		}
	}
	return ret
}