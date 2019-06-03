package main

import (
	"os"
	"bufio"
	"log"
)

func load(path string, length int) *map[string]bool {
	
	file, err := os.Open(dictionaryPath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	dictionary := make(map[string][]bool)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		if len(word) == length {
			_,valid := dictionary[word]
			if !valid {
				dictionary[word] = false
			}
		}
	}

	return &dictionary
}