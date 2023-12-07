package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"unicode"
)

var data []string

type hand struct {
	bonus      int
	bid        int
	cardValues []int
}

type pair struct {
	value      int
	duplicates int
}

func main() {
	lines, err := readLines("C:\\Users\\Mr.Mystery 1.0\\Desktop\\GitHub\\AoC_2023\\Puzzle_7\\resource\\data.txt")

	if err != nil {
		log.Fatal(err)
	}
	data = lines

	fmt.Printf("Result of part 1 is: %d \n", part1())
	fmt.Printf("Result of part 2 is: %d \n", part2())
}

func part1() int {
	var result int = 0

	var hands []hand
	for _, line := range data {
		hands = append(hands, convertToHand(line))
	}
	hands = checkHandsForBonus(hands)

	hands = sortHandsByValue(hands)
	hands = sortHandsByBonus(hands)

	for index := range hands {
		result += hands[index].bid * (index + 1)
	}

	return result
}

func part2() int {
	var result int = 0

	var hands []hand
	for _, line := range data {
		var hand = convertToHand(line)
		for index, number := range hand.cardValues {
			if number == 11 {
				hand.cardValues[index] = -1
			}
		}
		hands = append(hands, hand)
	}
	hands = checkHandsForBonusPart2(hands)

	hands = sortHandsByValue(hands)
	hands = sortHandsByBonus(hands)

	for index := range hands {
		result += hands[index].bid * (index + 1)
	}

	return result
}

func sortHandsByBonus(hands []hand) []hand {
	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].bonus < hands[j].bonus
	})
	return hands
}

func sortHandsByValue(hands []hand) []hand {
	sort.SliceStable(hands, func(i, j int) bool {
		for index := 0; index < 5; index++ {
			var difference = hands[i].cardValues[index] - hands[j].cardValues[index]

			if difference == 0 {
				continue
			} else if difference < 0 {
				return true
			} else {
				return false
			}
		}
		return true
	})
	return hands
}

// Gives bonus points for the hands based on their cards. This bonus is first and most important for sorting
func checkHandsForBonus(hands []hand) []hand {
	for index, hand := range hands {
		var pairs []pair
		// Find pairs
		for _, value := range hand.cardValues {
			var foundDuplicate bool = false
			for index, pair := range pairs {
				if pair.value == value {
					pair.duplicates++
					pairs[index] = pair
					foundDuplicate = true
					break
				}
			}
			if !foundDuplicate {
				pairs = append(pairs, pair{value, 1})
			}
		}

		// Check which bonus does apply based on pairs count and duplicates
		if len(pairs) == 1 {
			hand.bonus = 7
		} else if len(pairs) == 2 {
			var duplicates = pairs[0].duplicates
			if duplicates == 1 || duplicates == 4 {
				hand.bonus = 6
			} else {
				hand.bonus = 5
			}
		} else if len(pairs) == 3 {
			if pairs[0].duplicates == 3 || pairs[1].duplicates == 3 || pairs[2].duplicates == 3 {
				hand.bonus = 4
			} else {
				hand.bonus = 3
			}
		} else if len(pairs) == 4 {
			hand.bonus = 2
		} else if len(pairs) == 5 {
			hand.bonus = 1
		}
		hands[index] = hand
	}
	return hands
}

func checkHandsForBonusPart2(hands []hand) []hand {
	for index, hand := range hands {
		var pairs []pair
		var jokers int = 0

		// Find pairs
		for _, value := range hand.cardValues {
			if value == -1 {
				jokers++
				continue
			}

			var foundDuplicate bool = false
			for index, pair := range pairs {
				if pair.value == value {
					pair.duplicates++
					pairs[index] = pair
					foundDuplicate = true
					break
				}
			}
			if !foundDuplicate {
				pairs = append(pairs, pair{value, 1})
			}
		}
		if jokers == 5 {
			pairs = append(pairs, pair{2, 5})
		} else {
			var heighestPair pair = pairs[0]
			var heighestPairIndex int = 0
			for index, pair := range pairs {
				if pair.duplicates > heighestPair.duplicates {
					heighestPair = pair
					heighestPairIndex = index
				}
			}

			var betterpair = pairs[heighestPairIndex]
			betterpair.duplicates += jokers
			pairs[heighestPairIndex] = betterpair
		}

		// Check which bonus does apply based on pairs count and duplicates
		if len(pairs) == 1 {
			hand.bonus = 7
		} else if len(pairs) == 2 {
			var duplicates = pairs[0].duplicates
			if duplicates == 1 || duplicates == 4 {
				hand.bonus = 6
			} else {
				hand.bonus = 5
			}
		} else if len(pairs) == 3 {
			if pairs[0].duplicates == 3 || pairs[1].duplicates == 3 || pairs[2].duplicates == 3 {
				hand.bonus = 4
			} else {
				hand.bonus = 3
			}
		} else if len(pairs) == 4 {
			hand.bonus = 2
		} else if len(pairs) == 5 {
			hand.bonus = 1
		}
		hands[index] = hand
	}
	return hands
}

// Converts to a hand structure, replacing letter values with actual number values
func convertToHand(line string) hand {
	var hand hand

	var forBid bool

	for _, char := range line {
		if unicode.IsNumber(char) {
			if forBid {
				hand.bid = hand.bid*10 + int(char-'0')
			} else {
				hand.cardValues = append(hand.cardValues, int(char-'0'))
			}
		} else if unicode.IsLetter(char) {
			if char == 'A' {
				hand.cardValues = append(hand.cardValues, 14)
			} else if char == 'K' {
				hand.cardValues = append(hand.cardValues, 13)
			} else if char == 'Q' {
				hand.cardValues = append(hand.cardValues, 12)
			} else if char == 'J' {
				hand.cardValues = append(hand.cardValues, 11)
			} else if char == 'T' {
				hand.cardValues = append(hand.cardValues, 10)
			}
		} else {
			forBid = true
		}
	}
	return hand
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
