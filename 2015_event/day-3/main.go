package main

import (
	"fmt"
	"log"
	"os"
)

type Pos struct {
	x int
	y int
}

func main() {
	var data, err = os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sx, sy := 0, 0 // Santa
	rx, ry := 0, 0 // Robot

	visited := map[Pos]bool{}
	visited[Pos{0, 0}] = true

	for i, c := range data {
		if i%2 == 0 {
			// Santa
			switch c {
			case '^':
				sy++
			case 'v':
				sy--
			case '<':
				sx--
			case '>':
				sx++
			}
			visited[Pos{sx, sy}] = true
		} else {
			// Robot
			switch c {
			case '^':
				ry++
			case 'v':
				ry--
			case '<':
				rx--
			case '>':
				rx++
			}
			visited[Pos{rx, ry}] = true
		}
	}

	fmt.Println(len(visited))
}
