package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type dir struct {
	name string
	size int
}

func main() {
	//grid := getInputFileLines("input-test.txt")
	moveSet := getInputFileLines("input.txt")
	for _, m := range moveSet {
		log.Println(m)
	}
}

type Move struct {
	dir string
	num int
}

func getInputFileLines(fileName string) []Move {
	moveSet := make([]Move, 0)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		m := strings.Split(sc.Text(), " ")

		v, _ := strconv.Atoi(string(m[1]))
		curMove := Move{dir: m[0], num: v}
		moveSet = append(moveSet, curMove)
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return moveSet
}
