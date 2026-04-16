package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len (os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output. txt")
	return
    }

    inputFile := os.Args[1]
    outputFile := os.Args[2]

    content, err := os.ReadFile(inputFile)
    if err != nil {
	fmt.Println(err)
	return 
    }

    result := processText(string(content))

    err = os.WriteFile(outputFile, []byte(result), 0644)
    if err != nil {
	    fmt.Println(err)
	    return
    }
}

func processText(text string) string {
	words := strings.Fields(text)

	words = handleHexBin(words)
	

	return strings.Join(words, " ")

}

func handleHexBin(words []string) []string {
	result := []string {}

	for j := 0 ; j < len(words); j ++ {
		if words [j] == "(hex)" && len(result) > 0 {
			prev := result [len(result)-1]

			value, err := strconv.ParseInt(prev, 16, 64)
			if err == nil {
				result[len(result)-1] = fmt.Sprintf("%d", value)
			}
			continue
		}

	if words [j] == "(bin)" && len(result) > 0 {
		prev := result [len(result)-1]

		value, err := strconv.ParseInt(prev, 2, 64)
		if err == nil {
			result[len(result)-1] = fmt.Sprintf("%d", value)
		}
		continue
	}
	
	result = append(result, words[j])

	}

	return result
}
