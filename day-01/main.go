package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var inputFile = "input.txt"

func main() {
	n := 3
	caloriesPerElf := getCaloriesPerElf()
	topN := getTopN(caloriesPerElf, n)
	fmt.Printf("Top %v calories are:\n", n)
	for i, v := range topN {
		fmt.Printf("%v: %v\n", i, v)
	}
}

func getCaloriesPerElf() []int {
	var list = []int{}
	current := 0
	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	stringContent := strings.Split(string(content), "\n")
	for _, line := range stringContent {
		if line == "" {
			list = append(list, current)
			current = 0
		} else {
			i, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}

			current += i
		}
	}

	return list
}

func getTopN(slice []int, n int) []int {
	if n >= len(slice) {
		return slice
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
	return slice[:n]
}
