package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	//"slices"
	"strconv"
	//"unicode"
)

var data []string

func main() {
	lines, err := readLines("C:\\Users\\Mr.Mystery 1.0\\Desktop\\GitHub\\AoC_2023\\Puzzle_6\\resource\\data.txt")

	if err != nil {
		log.Fatal(err)
	}
	data = lines

	fmt.Printf("Result of part 1 is: %d \n", part1())
	fmt.Printf("Result of part 2 is: %d \n", part2())
}

func part1() int {
	var result int = 1

	var times = convertToIntSlice(strings.Fields(data[0])[1:])
	var bestT = convertToIntSlice(strings.Fields(data[1])[1:])

	for index := range times {
		var count = 0
		for num := 1; num < times[index]; num++ {
			if num*(times[index]-num) > bestT[index] {
				count++
			}
		}
		result *= count
	}

	return result
}

func part2() int {
	var result int = 0

	var times = strings.Fields(data[0])[1:]
	var bestT = strings.Fields(data[1])[1:]

	var givenTimeS string
	var bestScoreS string
	for _, part := range times {
		givenTimeS += part
	}
	for _, part := range bestT {
		bestScoreS += part
	}
	givenTime, _ := strconv.Atoi(givenTimeS)
	bestScore, _ := strconv.Atoi(bestScoreS)

	for num := 1; num < givenTime; num++ {
		if num*(givenTime-num) > bestScore {
			result++
		}
	}

	return result
}

func convertToIntSlice(stringSlice []string) []int {
	var intSlice []int
	for _, item := range stringSlice {
		var value, _ = strconv.Atoi(item)
		intSlice = append(intSlice, value)
	}
	return intSlice
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
