package main

import (
	"strings"
)

func bfs (dict *map[string]bool, start string, end string) []string {
	var queue []string
	var path []string


	dict[start] = true

	queue = append(queue,start)
	path = append(queue, start)

	for len(queue) > 0 {
		// Get value now so it doesn't change
		inQueue = len(queue)

		// Iterate over all in the queue at the minute
		for i := 0 ; i < inQueue; i++ {
			word := queue[0]
			queue = queue[1:] // 1 to end, so trim off the first

			dict[word] = true

			next := neighbours(dict, word)

			for n := range next {
				if n != end {
					queue = append(queue,n)
				} else {
					
				}
				
			}
			
		}
		
	}
}