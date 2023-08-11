package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var inputFile = "input.txt"

func main() {
	maxCalorieElf := getMaxCalories()
	fmt.Printf("Most calories: %v\n", maxCalorieElf)

	top3CalorieElfs := getTopThreeCalories()
	fmt.Printf("Top 3 calories: %v %v %v\n", top3CalorieElfs[0], top3CalorieElfs[1], top3CalorieElfs[2])
	fmt.Printf("Top 3 total calories: %v\n", top3CalorieElfs[0]+top3CalorieElfs[1]+top3CalorieElfs[2])
}

func getTopThreeCalories() []int {
	current := 0
	top3 := []int{0, 0, 0}

	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	stringContent := strings.Split(string(content), "\n")
	for _, line := range stringContent {
		if line == "" {
			indexOfSmallest := indexOfSmallestInt(top3)
			if current > top3[indexOfSmallest] {
				top3[indexOfSmallest] = current
			}
			current = 0
			continue
		}

		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		current += i
	}

	return top3
}

func indexOfSmallestInt(array []int) int {
	idx := 0
	min := math.MaxInt
	for i, v := range array {
		if v < min {
			min = v
			idx = i
		}
	}
	return idx
}

func getMaxCalories() int {
	var content, err = os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	max := 0
	current := 0
	stringContent := strings.Split(string(content), "\n")
	for _, line := range stringContent {
		if line == "" {
			if current > max {
				max = current
			}
			current = 0
			continue
		}

		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		current += i
	}

	return max
}
