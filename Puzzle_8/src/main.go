package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var data []string

type pair struct {
	left  string
	right string
}

func main() {
	lines, err := readLines("C:\\Users\\Mr.Mystery 1.0\\Desktop\\GitHub\\AoC_2023\\Puzzle_8\\resource\\data.txt")

	if err != nil {
		log.Fatal(err)
	}
	data = lines

	fmt.Printf("Result of part 1 is: %d \n", part1())
	fmt.Printf("Result of part 2 is: %d \n", part2())
}

func part1() int {
	var result int = 0

	var instructions string = data[0]
	var hashMap = getMap()

	var instructionIndex int = len(instructions) - 1
	var currentKey string = "AAA"
	for currentKey != "ZZZ" {
		if instructionIndex == len(instructions)-1 {
			instructionIndex = 0
		} else {
			instructionIndex++
		}
		if instructions[instructionIndex] == 'L' {
			currentKey = hashMap[currentKey].left
		} else {
			currentKey = hashMap[currentKey].right
		}
		result++
	}

	return result
}

func part2() int {
	var instructions string = data[0]
	var hashMap = getMap()

	var currentKeys []string

	for k := range hashMap {
		if k[2] == 'A' {
			currentKeys = append(currentKeys, k)
		}
	}

	var neededSteps []int
	for index := range currentKeys {
		var instructionIndex int = len(instructions) - 1
		var steps int
		for currentKeys[index][2] != 'Z' {
			if instructionIndex == len(instructions)-1 {
				instructionIndex = 0
			} else {
				instructionIndex++
			}

			if instructions[instructionIndex] == 'L' {
				currentKeys[index] = hashMap[currentKeys[index]].left
			} else {
				currentKeys[index] = hashMap[currentKeys[index]].right
			}
			steps++
		}
		neededSteps = append(neededSteps, steps)
	}

	return lcm(neededSteps...)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

func lcm(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = (res * nums[i]) / gcd(res, nums[i])
	}

	return res
}

func getMap() map[string]pair {
	var hashMap map[string]pair = make(map[string]pair, len(data)-2)

	for index := 2; index < len(data); index++ {
		var key, values, _ = strings.Cut(strings.Replace(strings.Replace(strings.Replace(data[index], "(", "", -1), ")", "", -1), " ", "", -1), "=")
		var leftAndRight = strings.Split(values, ",")
		hashMap[key] = pair{leftAndRight[0], leftAndRight[1]}
	}
	return hashMap
}

func readLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
