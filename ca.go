package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 5 || os.Args[1] != "-f" || os.Args[3] != "-keywords" {
		fmt.Println("Usage: go run ca.go -f <file_path> -keywords <keyword1,keyword2,...>")
		return
	}

	// Get the file path and keywords from the command line
	filePath := os.Args[2]
	keywords := strings.Split(os.Args[4], ",")

	// Load the txt file
	content, err := loadFile(filePath)
	if err != nil {
		fmt.Println("Error loading file:", err)
		return
	}

	// Analyze the file
	totalWords, wordOccurrences, wordCountByLength := analyzeFile(content, keywords)

	// Calculate word percentages
	wordPercentages := calculateWordPercentages(wordOccurrences, totalWords)

	// Display the results
	fmt.Printf("Total words in the file: %d\n", totalWords)

	fmt.Println("Word percentages based on the total number of strings loaded:")
	for word, percentage := range wordPercentages {
		fmt.Printf("  %s: %.2f%%\n", word, percentage)
	}

	fmt.Println("Occurrences of words or their substrings in the keyword list:")
	for word, count := range wordOccurrences {
		fmt.Printf("  %s: %d\n", word, count)
	}

	fmt.Printf("Word count by number of characters:\n")
	printWordCountByLength(wordCountByLength)

	// Some statistics ideas came from https://github.com/HynekPetrak/passat
	// Enumerate the 10 most used character sequences

	topSequences, sequenceCounts := enumerateTopSequences(content, 10, totalWords)
	fmt.Println("\nTop 10 most used character sequences:")
	for i, sequence := range topSequences {
		percentage := float64(sequenceCounts[sequence]) / float64(totalWords) * 100.0
		fmt.Printf("%d. %s (Sequence Count: %d, Percentage: %.2f%%)\n", i+1, sequence, sequenceCounts[sequence], percentage)
	}

	// Enumerate the top 10 most frequent alpha characters
	topAlphaChars, alphaCharCounts := enumerateTopChars(content, unicode.IsLetter, 10, totalWords)
	fmt.Println("\nTop 10 most frequent alpha characters:")
	printTopChars(topAlphaChars, alphaCharCounts, totalWords)

	// Enumerate the top 10 most frequent number characters
	topNumberChars, numberCharCounts := enumerateTopChars(content, unicode.IsNumber, 10, totalWords)
	fmt.Println("\nTop 10 most frequent number characters:")
	printTopChars(topNumberChars, numberCharCounts, totalWords)

	// Enumerate the top 10 most frequent symbol characters
	topSymbolChars, symbolCharCounts := enumerateTopChars(content, func(r rune) bool {
		return unicode.IsSymbol(r) || unicode.IsPunct(r)
	}, 10, totalWords)
	fmt.Println("\nTop 10 most frequent symbol characters:")
	printTopChars(topSymbolChars, symbolCharCounts, totalWords)

	// Additional rules
	fmt.Println("\nAdditional Rules:")
	ruleCountAndPercentage(content, "Four digits at the end", func(s string) bool {
		return len(s) >= 4 && unicode.IsDigit([]rune(s)[len(s)-4])
	}, totalWords)
	ruleCountAndPercentage(content, "Two digits at the end", func(s string) bool {
		return len(s) >= 2 && unicode.IsDigit([]rune(s)[len(s)-2])
	}, totalWords)
	ruleCountAndPercentage(content, "Upper + lower + num + symbol", func(s string) bool {
		return containsUppercase(s) && containsLowercase(s) && containsDigit(s) && containsSymbol(s)
	}, totalWords)
	ruleCountAndPercentage(content, "Three digits at the end", func(s string) bool {
		return len(s) >= 3 && unicode.IsDigit([]rune(s)[len(s)-3])
	}, totalWords)
	ruleCountAndPercentage(content, "First capital, last symbol", func(s string) bool {
		return len(s) > 0 && unicode.IsUpper([]rune(s)[0]) && containsSymbol(s[len(s)-1:])
	}, totalWords)
	ruleCountAndPercentage(content, "First capital, last number", func(s string) bool {
		return len(s) > 0 && unicode.IsUpper([]rune(s)[0]) && unicode.IsDigit([]rune(s)[len(s)-1])
	}, totalWords)
	ruleCountAndPercentage(content, "Digits are between 2000 and 2999", func(s string) bool {
		return isBetweenDigits(s, 2000, 2999)
	}, totalWords)
	ruleCountAndPercentage(content, "Digits are between 1900 and 1999", func(s string) bool {
		return isBetweenDigits(s, 1900, 1999)
	}, totalWords)
}

func loadFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + " "
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content, nil
}

func analyzeFile(content string, keywords []string) (int, map[string]int, map[int]int) {
	words := strings.Fields(content)
	totalWords := len(words)
	wordOccurrences := make(map[string]int)
	wordCountByLength := make(map[int]int)

	for _, word := range words {
		// Count occurrences or substring matches in the keyword list
		for _, targetWord := range keywords {
			if strings.Contains(strings.ToLower(word), strings.ToLower(targetWord)) {
				wordOccurrences[targetWord]++
				break
			}
		}

		// Count words by the number of characters
		wordCountByLength[len(word)]++
	}

	return totalWords, wordOccurrences, wordCountByLength
}

func calculateWordPercentages(wordOccurrences map[string]int, totalWords int) map[string]float64 {
	wordPercentages := make(map[string]float64)

	for word, count := range wordOccurrences {
		wordPercentages[word] = float64(count) / float64(totalWords) * 100.0
	}

	return wordPercentages
}

func printWordCountByLength(wordCountByLength map[int]int) {
	for length, count := range wordCountByLength {
		fmt.Printf("  Words with %d characters: %d\n", length, count)
	}
}

func enumerateTopSequences(content string, num int, totalWords int) ([]string, map[string]int) {
	sequences := make(map[string]int)
	for i := 0; i <= len(content)-num; i++ {
		sequence := content[i : i+num]
		sequences[sequence]++
	}

	var topSequences []string
	for sequence, count := range sequences {
		if count > 1 { // Adjust the threshold as needed
			topSequences = append(topSequences, sequence)
		}
	}

	sort.Slice(topSequences, func(i, j int) bool {
		return sequences[topSequences[i]] > sequences[topSequences[j]]
	})

	if len(topSequences) > num {
		topSequences = topSequences[:num]
	}

	return topSequences, sequences
}

func enumerateTopChars(content string, isCharType func(r rune) bool, num int, totalWords int) ([]string, map[string]int) {
	characters := make(map[string]int)
	for _, char := range content {
		if isCharType(char) {
			characters[string(char)]++
		}
	}

	var topChars []string
	for char := range characters {
		topChars = append(topChars, char)
	}

	sort.Slice(topChars, func(i, j int) bool {
		return characters[topChars[i]] > characters[topChars[j]]
	})

	if len(topChars) > num {
		topChars = topChars[:num]
	}

	return topChars, characters
}

func printTopChars(topChars []string, charCounts map[string]int, totalWords int) {
	for i, char := range topChars {
		percentage := float64(charCounts[char]) / float64(totalWords) * 100.0
		fmt.Printf("%d. %s (Count: %d, Percentage: %.2f%%)\n", i+1, char, charCounts[char], percentage)
	}
}

func containsUppercase(s string) bool {
	for _, char := range s {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

func containsLowercase(s string) bool {
	for _, char := range s {
		if unicode.IsLower(char) {
			return true
		}
	}
	return false
}

func containsDigit(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func containsSymbol(s string) bool {
	for _, char := range s {
		if unicode.IsSymbol(char) || unicode.IsPunct(char) {
			return true
		}
	}
	return false
}

func isBetweenDigits(s string, start, end int) bool {
	digitRegex := regexp.MustCompile("\\d+")
	matches := digitRegex.FindAllString(s, -1)

	for _, match := range matches {
		if len(match) >= 4 { // Ensure the match has at least 4 digits
			num, err := strconv.Atoi(match)
			if err == nil && num >= start && num <= end {
				return true
			}
		}
	}

	return false
}

func ruleCountAndPercentage(content string, rule string, condition func(string) bool, totalWords int) {
	count := 0
	words := strings.Fields(content)
	for _, word := range words {
		if condition(word) {
			count++
		}
	}
	percentage := float64(count) / float64(totalWords) * 100.0
	fmt.Printf("%s: Count: %d, Percentage: %.2f%%\n", rule, count, percentage)
}
