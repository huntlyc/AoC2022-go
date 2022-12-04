package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type elfPair struct {
	firstElf  string
	secondElf string
}

func main() {
	//elfPairs := getInputFileLines("input-test.txt")
	elfPairs := getInputFileLines("input.txt")
	part1(elfPairs)
	part2(elfPairs)
}

func part1(elfPairs []elfPair) {
	total := 0

	for _, pair := range elfPairs {

		// figure out ranges from pair strs
		firstElfRange := strings.Split(pair.firstElf, "-")
		e1start, err := strconv.Atoi(firstElfRange[0])
		if err != nil {
			log.Fatal(err)
		}
		e1end, err := strconv.Atoi(firstElfRange[1])
		if err != nil {
			log.Fatal(err)
		}

		secondElfRange := strings.Split(pair.secondElf, "-")
		e2start, err := strconv.Atoi(secondElfRange[0])
		if err != nil {
			log.Fatal(err)
		}
		e2end, err := strconv.Atoi(secondElfRange[1])
		if err != nil {
			log.Fatal(err)
		}

		// compare elf ranges, if one pair includes the other, tally it up
		if (e1start <= e2start && e1end >= e2end) || (e2start <= e1start && e2end >= e1end) {
			total += 1
		}

	}

	log.Println("Part1 total: ", total)
}

func part2(elfPairs []elfPair) {
	total := 0

	for _, pair := range elfPairs {

		// figure out ranges from pair strs
		firstElfRange := strings.Split(pair.firstElf, "-")
		e1start, err := strconv.Atoi(firstElfRange[0])
		if err != nil {
			log.Fatal(err)
		}
		e1end, err := strconv.Atoi(firstElfRange[1])
		if err != nil {
			log.Fatal(err)
		}

		secondElfRange := strings.Split(pair.secondElf, "-")
		e2start, err := strconv.Atoi(secondElfRange[0])
		if err != nil {
			log.Fatal(err)
		}

		e2end, err := strconv.Atoi(secondElfRange[1])
		if err != nil {
			log.Fatal(err)
		}

		// check for any overlap
		if (e1start <= e2start && e1end >= e2start) ||
			(e1start <= e2end && e1end >= e2end) ||
			(e2start <= e1start && e2end >= e1start) ||
			(e2start <= e1end && e2end >= e1end) {
			total += 1
		}

	}

	log.Println("Part2 total: ", total)
}

func getInputFileLines(fileName string) []elfPair {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	allPairs := make([]elfPair, 0)

	for sc.Scan() {
		pairs := strings.Split(sc.Text(), ",")
		allPairs = append(allPairs, elfPair{pairs[0], pairs[1]})
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return allPairs
}
