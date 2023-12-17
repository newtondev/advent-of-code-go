package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var symbols = [][]bool{}
var engineParts [][]enginePart

// result stores the results for printing to output
type result struct {
	p1, p2 uint32
}

// digit represents a single digit in an engine part
type digit struct {
	row  int
	col  int
	raw  int
	pass bool
}

// enginePart represents a single engine part
type enginePart struct {
	digits []digit
}

// realValue returns the numeric valuee of all digits combined
func (e *enginePart) realValue() (int, error) {
	var numStr string
	for _, digit := range e.digits {
		numStr += fmt.Sprintf("%d", digit.raw)
	}
	n, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// checkDigit checks if any neighbouring characters are true (in symbol matrix)
func checkDigit(d digit) digit {
	i := d.col - 1
	for ; i < len(symbols[d.row]) && i <= d.col+1; i += 1 {
		if i < 0 {
			continue
		}
		if d.row != 0 {
			if symbols[d.row-1][i] {
				d.pass = true
				break
			}
		}
		if symbols[d.row][i] {
			d.pass = true
			break
		}
		if d.row != len(symbols)-1 {
			if symbols[d.row+1][i] {
				d.pass = true
				break
			}
		}
	}
	return d
}

// checkDigits runs each digit in an engine part through checkDigit()
func checkDigits(i, j int, line string) (int, int, error) {
	enginePart := enginePart{}
	for ; j < len(line); j += 1 {
		if !unicode.IsNumber(rune(line[j])) {
			break
		}
		enginePart.digits = append(enginePart.digits, checkDigit(digit{row: i, col: j, raw: int(line[j]) - '0', pass: false}))
	}

	if anyTrue(enginePart.digits) {
		engineParts[i] = append(engineParts[i], enginePart)
		n, err := enginePart.realValue()
		if err != nil {
			return 0, 0, err
		}
		return j, n, nil
	}
	return j, 0, nil
}

// symbolToBool generates a boolean matrix that represents the placement of symbols
func symbolToBool(line string) []bool {
	b := []bool{}
	for _, c := range line {
		b = append(b, !(unicode.IsNumber(c) || c == '.'))
	}
	return b
}

// solvePartOne deciphers which numbers represent engine parts
// it returns the sum of all engine parts
func solvePartOne(input string) (int, error) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		symbols = append(symbols, symbolToBool(line))
	}
	engineParts = make([][]enginePart, len(lines))

	sum := 0
	for i, line := range lines {
		for j := 0; j < len(line); j += 1 {
			if unicode.IsNumber(rune(line[j])) {
				next, n, err := checkDigits(i, j, line)
				if err != nil {
					return 0, err
				}
				sum += n
				j = next
			}
		}
	}

	return sum, nil
}

// getGearRatio returns the product of engine parts if exactly
// two neighbour a gear
func getGearRatio(row, col int) (int, error) {
	partsConsidered := []enginePart{}

	i := row - 1
	for ; i < len(engineParts) && i <= row+1; i += 1 {
		if i < 0 {
			continue
		}
		for _, enginePart := range engineParts[i] {
			for _, digit := range enginePart.digits {
				if digit.pass && digit.col == col-1 || digit.col == col || digit.col == col+1 {
					partsConsidered = append(partsConsidered, enginePart)
					break
				}
			}
		}
	}
	if len(partsConsidered) == 2 {
		n1, err := partsConsidered[0].realValue()
		if err != nil {
			return 0, err
		}
		n2, err := partsConsidered[1].realValue()
		if err != nil {
			return 0, err
		}
		return n1 * n2, nil
	}
	return 0, nil
}

func solvePartTwo(input string) (int, error) {
	lines := strings.Split(input, "\n")
	sum := 0

	for i, line := range lines {
		for j, c := range line {
			if rune(c) == '*' {
				n, err := getGearRatio(i, j)
				if err != nil {
					return 0, err
				}
				sum += n
			}
		}
	}

	return sum, nil
}

func main() {
	res := solve(readFileInput("inputs/03.txt"))

	fmt.Println(res.p1)
	fmt.Println(res.p2)
}

func solve(input string) result {
	p1, err := solvePartOne(input)
	if err != nil {
		panic(err)
	}

	p2, err := solvePartTwo(input)
	if err != nil {
		panic(err)
	}

	return result{p1: uint32(p1), p2: uint32(p2)}
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

func anyTrue(digits []digit) bool {
	for _, digit := range digits {
		if digit.pass {
			return true
		}
	}
	return false
}
