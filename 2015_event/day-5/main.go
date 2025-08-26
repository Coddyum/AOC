package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	niceCountPart1 := 0
	niceCountPart2 := 0

	for _, line := range lines {
		if isNicePart1(line) {
			niceCountPart1++
		}
		if isNicePart2(line) {
			niceCountPart2++
		}
	}

	fmt.Println("Part 1 - Nice strings:", niceCountPart1)
	fmt.Println("Part 2 - Nice strings:", niceCountPart2)
}

// RULES PART 1

func isNicePart1(s string) bool {
	return hasThreeVowels(s) && hasDoubleLetter(s) && !hasForbiddenSubstring(s)
}

func hasThreeVowels(s string) bool {
	vowels := "aeiou"
	count := 0
	for _, ch := range s {
		if strings.ContainsRune(vowels, ch) {
			count++
			if count >= 3 {
				return true
			}
		}
	}
	return false
}

func hasDoubleLetter(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func hasForbiddenSubstring(s string) bool {
	forbidden := []string{"ab", "cd", "pq", "xy"}
	for _, f := range forbidden {
		if strings.Contains(s, f) {
			return true
		}
	}
	return false
}

// RULES PART 2

func isNicePart2(s string) bool {
	return hasPairTwice(s) && hasRepeatWithGap(s)
}

func hasPairTwice(s string) bool {
	pairs := make(map[string]int)
	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		if idx, ok := pairs[pair]; ok {
			if i-idx >= 2 {
				return true
			}
		} else {
			pairs[pair] = i
		}
	}
	return false
}

func hasRepeatWithGap(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}
