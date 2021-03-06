// --- Day 2: Dive! ---

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func day2() (uint64, uint64) {
	file, err := os.Open("../inputs/2.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var pos, aim, dth uint64 = 0, 0, 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		op := line[0]
		val, err := strconv.ParseUint(line[1], 10, 64)

		if err != nil {
			panic(err)
		}

		switch op {
		case "down":
			aim += val
		case "up":
			aim -= val
		case "forward":
			pos += val
			dth += aim * val
		}
	}

	return pos * aim, pos * dth
}
