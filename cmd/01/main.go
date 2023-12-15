package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Pair struct {
	First uint32
	Last  uint32
}

type PairInterface interface {
	Result() uint32
}

func (p Pair) Result() uint32 {
	return p.First*10 + p.Last
}

func main() {
	partOneTotal := solvePartOne(readFileInput("inputs/01.txt"))
	fmt.Printf("%d\n", partOneTotal)

	partTwoTotal := solvePartTwo(readFileInput("inputs/01.txt"))
	fmt.Printf("%d\n", partTwoTotal)
}

func solvePartOne(input string) uint32 {
	lines := strings.Split(input, "\n")

	total := uint32(0)
	for _, line := range lines {
		digits := make([]uint32, 0)
		for _, c := range line {
			n, err := strconv.ParseUint(string(c), 10, 32)
			if err == nil {
				digits = append(digits, uint32(n))
			}
		}

		pair := Pair{0, 0}
		for _, n := range digits {
			switch n {
			case 0:
				continue
			default:
				if pair.First == 0 {
					pair.First = n
					pair.Last = n
				} else {
					pair.Last = n
				}
			}
		}

		total += pair.Result()
	}

	return total
}

func solvePartTwo(input string) uint32 {
	lines := strings.Split(input, "\n")

	total := uint32(0)
	re := regexp.MustCompile(`(oneight|twone|threeight|fiveight|sevenine|eightwo|eighthree|nineight|one|two|three|four|five|six|seven|eight|nine|\d)`)

	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		digits := make([]uint32, 0)

		for _, m := range matches {
			switch m {
			case "one":
				digits = append(digits, 1)
			case "two":
				digits = append(digits, 2)
			case "three":
				digits = append(digits, 3)
			case "four":
				digits = append(digits, 4)
			case "five":
				digits = append(digits, 5)
			case "six":
				digits = append(digits, 6)
			case "seven":
				digits = append(digits, 7)
			case "eight":
				digits = append(digits, 8)
			case "nine":
				digits = append(digits, 9)
			case "oneight":
				digits = append(digits, 1, 8)
			case "twone":
				digits = append(digits, 2, 1)
			case "threeight":
				digits = append(digits, 3, 8)
			case "fiveight":
				digits = append(digits, 5, 8)
			case "sevenine":
				digits = append(digits, 7, 9)
			case "eightwo":
				digits = append(digits, 8, 2)
			case "eighthree":
				digits = append(digits, 8, 3)
			case "nineight":
				digits = append(digits, 9, 8)
			default:
				n, err := strconv.ParseUint(m, 10, 32)
				if err == nil {
					digits = append(digits, uint32(n))
				}
			}
		}

		pair := Pair{0, 0}
		for _, n := range digits {
			switch n {
			case 0:
				continue
			default:
				if pair.First == 0 {
					pair.First = n
					pair.Last = n
				} else {
					pair.Last = n
				}
			}
		}

		total += pair.Result()
	}

	return total
}

func readFileInput(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return string(data)
}
