package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var regex = map[string]*regexp.Regexp{
	"red":   regexp.MustCompile(`(?P<red>\d+) red`),
	"green": regexp.MustCompile(`(?P<green>\d+) green`),
	"blue":  regexp.MustCompile(`(?P<blue>\d+) blue`),
	"game":  regexp.MustCompile(`Game (?P<game>\d+)`),
}

var limits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func getIDForGame(identifier string) (int, error) {
	m := regex["game"].FindStringSubmatch(identifier)
	i := regex["game"].SubexpIndex("game")
	id, err := strconv.Atoi(m[i])
	if err != nil {
		return 0, err
	}
	return id, nil
}

func isValidGame(data string) (bool, error) {
	for _, colour := range []string{"red", "green", "blue"} {
		m := regex[colour].FindAllStringSubmatch(data, -1)
		for _, s := range m {
			n, err := strconv.Atoi(s[1])
			if err != nil {
				return false, err
			}
			if n > limits[colour] {
				return false, nil
			}
		}
	}
	return true, nil
}

func main() {
	res := solve(readFileInput("inputs/02.txt"))

	fmt.Println(res.p1)
	fmt.Println(res.p2)
}

type result struct {
	p1, p2 uint32
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

func solvePartOne(input string) (int, error) {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		identifier, data := func() (string, string) {
			x := strings.Split(line, ":")
			return x[0], x[1]
		}()

		isValidGame, err := isValidGame(data)
		if err != nil {
			return 0, err
		}
		if isValidGame {
			id, err := getIDForGame(identifier)
			if err != nil {
				return 0, err
			}
			sum += id
		}
	}

	return sum, nil
}

func powerForGame(data string) (int, error) {
	var counter = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, colour := range []string{"red", "green", "blue"} {
		m := regex[colour].FindAllStringSubmatch(data, -1)
		for _, s := range m {
			n, err := strconv.Atoi(s[1])
			if err != nil {
				return 0, err
			}
			if counter[colour] < n {
				counter[colour] = n
			}
		}
	}
	return counter["red"] * counter["green"] * counter["blue"], nil
}

func solvePartTwo(input string) (int, error) {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		_, data := func() (string, string) {
			x := strings.Split(line, ":")
			return x[0], x[1]
		}()

		pow, err := powerForGame(data)
		if err != nil {
			return 0, err
		}
		sum += pow
	}

	return sum, nil
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
