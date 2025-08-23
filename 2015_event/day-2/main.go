package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var data, err = os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var chars = strings.Split(string(data), "\n")

	total := 0
	totalRibbon := 0
	for _, line := range chars {
		parts := strings.Split(line, "x")

		// strings parts to int
		l, _ := strconv.Atoi(parts[0])
		w, _ := strconv.Atoi(parts[1])
		h, _ := strconv.Atoi(parts[2])

		lw := l * w
		wh := w * h
		hl := h * l

		surface := 2*lw + 2*wh + 2*hl
		minSurface := min(lw, wh, hl)

		dims := []int{l, w, h}
		sort.Ints(dims)

		small1 := dims[0]
		small2 := dims[1]

		RibbonTour := 2 * (small1 + small2)
		RibbonArc := l * w * h

		totalRibbon += RibbonTour + RibbonArc
		total += surface + minSurface
	}

	fmt.Println("Total:", total)
	fmt.Println("Total Ribbon:", totalRibbon)
}
