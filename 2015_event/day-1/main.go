package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var data, err = os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	floor := 0
	for i, c := range string(data) {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor == -1 {
			answer := i + 1
			fmt.Println("Answer:", answer)
			break
		}
	}

	fmt.Println("Floor:", floor)
}
