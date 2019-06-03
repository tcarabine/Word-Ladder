package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
)

func main () {

	dictionaryPath := "/workspaces/wordchain/data/dictionary.txt"
	startWord := "cat"
	//endWord := "dog"
	length := len(startWord)

	dictionary := load(dictionaryPath)

	allWords := make([]string, 0, len(dictionary))
	for key := range dictionary {
		allWords = append(allWords, key)
	}

	

}