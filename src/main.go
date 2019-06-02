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

	file, err := os.Open(dictionaryPath)
	if err != nil {
		log.Fatal(err)
	}

	dictionary := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) == length {
			_,valid := dictionary[word]
			if !valid {
				dictionary[word] = []string{}
			}
		}
	}

	allWords := make([]string, 0, len(dictionary))
	for key := range dictionary {
		allWords = append(allWords, key)
	}

	for w := range allWords {
		wSlice := strings.Split(allWords[w], "")

		for k := range allWords {
			if k != w {
				kSlice := strings.Split(allWords[k], "")
				diff := 0

				for i := 0; i < length; i++ {
					if kSlice[i] != wSlice[i] {
						diff++
					}
				}
				if diff == 1 {
					dictionary[allWords[k]] = append(dictionary[allWords[k]],allWords[w])

					fmt.Println( "Added " ,allWords[w], " to ", allWords[k])
				}
			}
		}
	}

}