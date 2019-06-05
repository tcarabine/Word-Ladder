package main

func bfs (dict map[string]node, start string, end string) []string {
	var queue []node
	var path []string

	queue = append(queue,dict[start])
	path = append(path, end)

	for len(queue) > 0 {
		// Get value now so it doesn't change
		inQueue := len(queue)

		// Iterate over all in the queue at the minute
		for i := 0 ; i < inQueue; i++ {
			vertex := queue[0]
			queue = queue[1:] // 1 to end, so trim off the first

			dict[vertex.word] = vertex

			if vertex.word == end {
				x := dict[end].parent

				for x != start {
					path = append(path, x)
					x = dict[x].parent
				}
				path = append(path, start)
			} else {
				next := neighbours(dict, vertex.word)
				for n := range next {
					queue = append(queue,dict[next[n]])
				}
			}
		}
	}

	return path
}