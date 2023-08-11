package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var inputFile = "input.txt"

func main() {
	maxCalorieElf := getMaxCalories()
	fmt.Printf("Most calories: %v\n", maxCalorieElf)
}

func getMaxCalories() int {
	var content, err = os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var max = 0
	var current = 0
	var stringContent = strings.Split(string(content), "\n")
	for _, line := range stringContent {
		if line == "" {
			if current > max {
				max = current
			}
			current = 0
			continue
		}

		var i, err = strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		current += i
	}

	return max
}
