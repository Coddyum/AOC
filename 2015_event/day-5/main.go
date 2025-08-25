package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Get Strings
	var data, err = os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// init strings
	var chars = strings.Split(string(data), "\n")

	// Tableau step by step
	voyellesString := []string{}
	twiceString := []string{}
	niceString := []string{}

	// check si la chaine contient au moins 3 fois
	for _, c := range chars {
		if minVoyelles(c) {
			voyellesString = append(voyellesString, c)
		}
	}
	// Check si un lettre est doubler
	for _, v := range voyellesString {
		if containsDoubleLetter(v) {
			twiceString = append(twiceString, v)
		}
	}
	// check si un la strings contiant des chaines intedite
	for _, f := range twiceString {
		if !hasFobiddenSubString(f) {
			niceString = append(niceString, f)
		}
	}

	fmt.Println("voyellesString: ", len(voyellesString))
	fmt.Println("TwiceString: ", len(twiceString))
	fmt.Println("NiceString : ", len(niceString))
}

// Check si la string contient au minimum 3 voyelles parmis la liste suivante (a,e,i,o,u)
func minVoyelles(t string) bool {
	voyelles := "aeiou"
	count := 0

	for i := 0; i < len(t); i++ {
		if strings.ContainsRune(voyelles, rune(t[i])) {
			count++
			if count >= 3 {
				return true
			}
		}
	}
	return false
}

type PairsString struct {
	t1 string
	t2 string
}

// Check si la string contient un lettre en double (e.g : aa / ee / bb)
func containsDoubleLetter(t string) bool {
	t1 := ""
	t2 := ""
	pairs := map[PairsString]string{}
	for i := 0; i < len(t)-1; i++ {
		if t[i] == t[i+1] {
			t1 = string(t[i])
			t2 = string(t[i+1])
			pairs[PairsString{t1, t2}] = ""
		}
		if len(pairs) >= 1 {
			fmt.Println(pairs)
		}
	}
	return false
}

// Check si la string contient des charactère interdis
func hasFobiddenSubString(t string) bool {
	forbidden := []string{"ab", "cd", "pq", "xy"}
	for _, f := range forbidden {
		if strings.Contains(t, f) {
			return true
		}
	}
	return false
}
