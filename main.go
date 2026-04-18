package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// main reads input file, processes the text,
// and writes the transformed output to a file.

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

// processText is the main transformation pipeline.
// It applies grouped punctuation fixes, splits text into words,
// applies hex/bin conversion, case transformations, article correction,
// then rebuilds the string and fixes punctuation and quotes.

func processText(text string) string {
	text = fixGroupedPunct(text)

	words := strings.Fields(text)

	words = handleHexBin(words)
	words = handleCase(words)
	words = fixArticles(words)

	text = strings.Join(words, " ")

	text = handlePunctuation(text)
	text = fixQuotes(text)

	return text

}

// handleHexBin scans for (hex) and (bin) tokens.
// It converts the previous word from hexadecimal or binary
// into its decimal representation.

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

// handleCase processes case transformation tags:
// (up), (low), (cap) and multi-word versions like (up, 2).
// It modifies the previously added words accordingly.

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

// fixGroupedPunct handles grouped punctuation such as
// "..." and "!?" so they remain attached correctly
// before further punctuation processing.

func fixGroupedPunct(text string) string {
	text = strings.ReplaceAll(text, " ...", "...")
	text = strings.ReplaceAll(text, "...", "...")
	text = strings.ReplaceAll(text, " !?", "!?")
	return text
}

// fixArticles converts "a" to "an" when the next word
// begins with a vowel or 'h', according to project rules.

func fixArticles(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" {
			next := strings.ToLower(words[i+1])
			if len(next) > 0 && strings.ContainsRune("aeiouh", rune(next[0])) {
				words[i] = "an"
			}
		}
	}
	return words
}

// handlePunctuation removes spaces before punctuation
// characters and ensures correct spacing after them.

func handlePunctuation(text string) string {
	text = strings.ReplaceAll(text, " ,", ",")
	text = strings.ReplaceAll(text, " .", ".")
	text = strings.ReplaceAll(text, " !", "!")
	text = strings.ReplaceAll(text, " ?", "?")
	text = strings.ReplaceAll(text, " :", ":")
	text = strings.ReplaceAll(text, " ;", ";")

	return text
}

// fixQuotes removes unnecessary spaces inside single quotes
// and ensures quoted text is formatted correctly.

func fixQuotes(text string) string {
	text = strings.ReplaceAll(text, " ' ", "'")
	text = strings.ReplaceAll(text, "' ", "'")
	text = strings.ReplaceAll(text, " '", "'")
	return text
}

