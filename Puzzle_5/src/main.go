package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

var data []string

func main() {
	lines, err := readLines("C:\\Users\\Mr.Mystery 1.0\\Desktop\\GitHub\\AoC_2023\\Puzzle_5\\resource\\data.txt")

	if err != nil {
		log.Fatal(err)
	}
	data = lines

	fmt.Printf("Result of part 1 is: %d \n", part1())
	fmt.Printf("Result of part 2 is: %d \n", part2())
}

func part1() int {
	seeds := getSeeds()
	digits := getDigits()

	for seedIndex := range seeds {
		for section := 0; section < len(digits); section++ {
			for index := 0; index < len(digits[section]); index += 3 {
				var source = digits[section][index+1]
				var destination = digits[section][index]
				var length = digits[section][index+2]

				if source <= seeds[seedIndex] && source+length >= seeds[seedIndex] {
					var difference = seeds[seedIndex] - source

					seeds[seedIndex] = destination + difference
					break
				}
			}
		}
	}

	return slices.Min(seeds)
}

func part2() int {
	digits := getDigits()

	var smallestSeed = 2099999999
	var seedRanges = strings.Split(data[0], " ")

	for index := 1; index < len(seedRanges); index += 2 {
		value, err := strconv.Atoi(seedRanges[index])
		value2, err2 := strconv.Atoi(seedRanges[index+1])
		if err != nil && err2 != nil {
			continue
		}
		for rangeIndex := 0; rangeIndex < value2; rangeIndex++ {
			var seed = value + rangeIndex

			for section := 0; section < len(digits); section++ {
				for index := 0; index < len(digits[section]); index += 3 {
					var source = digits[section][index+1]
					var destination = digits[section][index]
					var length = digits[section][index+2]

					if source <= seed && source+length >= seed {
						var difference = seed - source

						seed = destination + difference
						break
					}
				}
			}
			if seed < smallestSeed {
				smallestSeed = seed
			}
		}
	}
	return smallestSeed
}

func getSeeds() []int {
	var seeds []int

	for _, num := range strings.Split(data[0], " ") {
		value, err := strconv.Atoi(num)
		if err != nil {
			continue
		}
		seeds = append(seeds, value)
	}

	return seeds
}

func getDigits() [][]int {
	var digits [][]int = make([][]int, 7)
	var index int = -2

	for _, line := range data {
		if line == "" {
			continue
		}
		if unicode.IsLetter([]rune(line)[0]) {
			index++
			continue
		}
		var digitChars []string = strings.Split(line, " ")
		for _, digit := range digitChars {
			value, err := strconv.Atoi(digit)
			if err != nil {
				log.Fatal(err)
			}
			digits[index] = append(digits[index], value)
		}
	}
	return digits
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
