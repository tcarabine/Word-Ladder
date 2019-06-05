package main

import (
	"fmt"
	"time"
)

type node struct {
	parent string
	word string
	visited bool
}

func main () {
	tStart := time.Now()

	dictionaryPath := "/workspaces/Word-Ladder/data/dictionary.txt"
	startWord := "cat"
	endWord := "dog"
	length := len(startWord)

	dictionary := load(dictionaryPath, length)

	tElapsed := time.Since(tStart)

	fmt.Println("Loaded dictionary in ",tElapsed.String())

	tStartBFS := time.Now()
	path := bfs(dictionary,startWord,endWord)

	fmt.Println("Got there in ", len(path))
	fmt.Printf("%v\n", path)
	tElapsedBFS := time.Since(tStartBFS)
	fmt.Println("Found path in ",tElapsedBFS.String())
	tTotal := time.Since(tStart)
	fmt.Println("Total run time ", tTotal.String())
}