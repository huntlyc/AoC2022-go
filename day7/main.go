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
	// 	hist := getInputFileLines("input-test.txt")
	hist := getInputFileLines("input.txt")
	dirList := getUnorderedDirSizesFromHistory(hist)

	part1Sum := 0
	for _, d := range dirList {
		if d.size <= 100_000 {
			part1Sum += d.size
		}
	}
	log.Println("p1 ", part1Sum)

	hdSize := 70_000_000
	updateSize := 30_000_000
	freeSpace := hdSize - dirList[0].size

	smallest := dirList[0] // start with root ;)
	for _, d := range dirList {
		if freeSpace+d.size >= updateSize {
			if d.size < smallest.size {
				smallest = d
			}
		}
	}

	log.Println("p2 ", smallest)
}

func getUnorderedDirSizesFromHistory(hist []string) []dir {

	// list of names/sizes
	tmpDirSizeList := make([]dir, 0)
	finalDirSizeList := make([]dir, 0) // final list will include "/"

	// vars to keep track of things while parsing history
	dirStack := make([]string, 0)
	dirSizeStack := make([]int, 0)
	isListingFiles := false
	rootSizeTally := 0

	for hi, e := range hist {
		entryParts := strings.Split(e, " ")

		if entryParts[0] == "$" { // cmd
			switch entryParts[1] {
			case "ls":
				isListingFiles = true
			case "cd":
				isListingFiles = false

				if entryParts[2] == ".." {
					lastdir := dir{
						name: dirStack[len(dirStack)-1],
						size: dirSizeStack[len(dirSizeStack)-1],
					}
					if lastdir.name != "/" {
						tmpDirSizeList = append(tmpDirSizeList, lastdir)
					}

					dirStack = dirStack[:len(dirStack)-1]
					dirSizeStack = dirSizeStack[:len(dirSizeStack)-1]
				} else { // new dir
					dirStack = append(dirStack, entryParts[2])
					dirSizeStack = append(dirSizeStack, 0)
				}
			}
		} else { //dir or file entry

			if isListingFiles {
				// if a file record its size
				if entryParts[0] != "dir" {
					fileSize, _ := strconv.Atoi(entryParts[0])
					for di := 0; di < len(dirSizeStack); di++ {
						dirSizeStack[di] += fileSize
					}

					rootSizeTally += fileSize

					// last item in history, process and then prepend root "/" to list
					if hi == len(hist)-1 {
						lastdir := dir{
							name: dirStack[len(dirStack)-1],
							size: dirSizeStack[len(dirSizeStack)-1],
						}

						if lastdir.name != "/" {
							tmpDirSizeList = append(tmpDirSizeList, lastdir)
						}

						rootDir := dir{
							name: "/",
							size: rootSizeTally,
						}

						finalDirSizeList = append(finalDirSizeList, rootDir)
						finalDirSizeList = append(finalDirSizeList, tmpDirSizeList...)
						tmpDirSizeList = finalDirSizeList
					}
				}

			}
		}
	}
	return finalDirSizeList
}

func getInputFileLines(fileName string) []string {
	streams := make([]string, 0)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		streams = append(streams, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return streams
}
