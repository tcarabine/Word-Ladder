package main

import (
	"fmt"
	"time"
	"flag"
)

type node struct {
	parent string
	word string
	visited bool
}

func main () {
	tStart := time.Now()
	
	//dictionaryPath := "/workspaces/Word-Ladder/data/dictionary.txt"
	dictionaryPath := "/workspaces/wordchain/data/dictionary.txt"

	var startWord,endWord string
	flag.StringVar(&startWord, "start", "sport", "Word to start search from")
	flag.StringVar(&endWord, "end", "crate", "Word to end on")

	flag.Parse()

	length := len(startWord)

	dictionary := load(dictionaryPath, length)

	tElapsed := time.Since(tStart)

	fmt.Println("Loaded dictionary in ",tElapsed.String())

	tStartBFS := time.Now()
	path := bfs(dictionary,startWord,endWord)

	fmt.Println("Got there in ", len(path))

	for i := 0; i < len(path)/2; i++ {
		j := len(path) - i - 1
		path[i], path[j] = path[j], path[i]
	}

	for p := range path {
		if p < len(path)-1 {
			fmt.Printf("%v -> ", path[p])
		} else {
			fmt.Printf("%v \n", path[p])
		}
	}

	tElapsedBFS := time.Since(tStartBFS)
	fmt.Println("Found path in ",tElapsedBFS.String())
	tTotal := time.Since(tStart)
	fmt.Println("Total run time ", tTotal.String())
}