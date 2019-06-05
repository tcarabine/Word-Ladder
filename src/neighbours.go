package main

import (
	"strings"
)

func neighbours (dict map[string]node, word string, end string) []string {
	var ret []string
	wordSlice := strings.Split(word, "")

	for compare := range dict {
		//if word != compare {
		if (word != compare) && !(dict[compare].visited) {

			compareSlice := strings.Split(compare, "")
			diff := 0

			for i := 0; i < len(word); i++ {
				if wordSlice[i] != compareSlice[i] {
					diff++
				}
			}
			//if (diff == 1) && !(dict[compare].visited) {
			if (diff == 1)  {

				// If there is only one letter different and the value in the dictionary is false - not visited, add to return slice
				ret = append(ret,compare)
				node := dict[compare]
				node.parent = word
				node.visited = true
				dict[compare] = node
				if compare == end {
					return append(make([]string,1),end)
				}
			}
		}
	}
	return ret
}