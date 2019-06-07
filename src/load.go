package main

import (
	"os"
	"bufio"
	"log"
)

func load(path string, length int) map[string]node {
	
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	dictionary := make(map[string]node)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		if len(word) == length {
			// _,valid := dictionary[word]
			// if !valid {
				dictionary[word] = node{"",word,false}
			//}
		}
	}

	return dictionary
}