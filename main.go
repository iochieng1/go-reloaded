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

    err = os.WriteFile(outputFile, []byte(result + "\n"), 0644)
    if err != nil {
	    fmt.Println(err)
	    return
    }
}

func processText(text string) string {
	words := strings.Fields(text)

	words = handleHexBin(words)
	words = handleCase(words)
	

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

func handleCase(words []string) []string {
	result := []string{}

	for j := 0; j < len(words); j++ {

		if words[j] == "(up)" && len(result) > 0 {
			result[len(result)-1] =
				strings.ToUpper(result[len(result)-1])
			continue
		}

		if words[j] == "(low)" && len(result) > 0 {
			result[len(result)-1] =
				strings.ToLower(result[len(result)-1])
			continue
		}

		if words[j] == "(cap)" && len(result) > 0 {
			word := result[len(result)-1]
			result[len(result)-1] =
				strings.ToUpper(string(word[0])) +
					strings.ToLower(word[1:])
			continue
		}

		if words[j] == "(up," && j+1 < len(words) {
			n, err := strconv.Atoi(strings.TrimRight(words[j+1], ")"))
			if err == nil {
				for k := len(result) - n; k < len(result); k++ {
					if k >= 0 {
						result[k] = strings.ToUpper(result[k])
					}
				}
			}
			j++
			continue
		}

		if words[j] == "(low," && j+1 < len(words) {
			n, err := strconv.Atoi(strings.TrimRight(words[j+1], ")"))
			if err == nil {
				for k := len(result) - n; k < len(result); k++ {
					if k >= 0 {
						result[k] = strings.ToLower(result[k])
					}
				}
			}
			j++
			continue
		}

		if words[j] == "(cap," && j+1 < len(words) {
			n, err := strconv.Atoi(strings.TrimRight(words[j+1], ")"))
			if err == nil {
				for k := len(result) - n; k < len(result); k++ {
					if k >= 0 {
						word := result[k]
						result[k] =
							strings.ToUpper(string(word[0])) +
								strings.ToLower(word[1:])
					}
				}
			}
			j++
			continue
		}

		result = append(result, words[j])
	}

	return result
}

func getCount(token string) int {
	token = strings.Trim(token, "()")
	parts := strings.Split(token, ",")

	if len(parts) == 2 {
		n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err == nil {
			return n
		}
	}
	return 1
}
