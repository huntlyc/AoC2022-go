package main

import (
	"bufio"
	"log"
	"os"
)


func main() {
	//streams := getInputFileLines("input-test.txt")
	streams := getInputFileLines("input.txt")
    for _, str := range streams {
        ans := getStreamStartIdx(str, 4)
        log.Println("Part1: ", ans)
        ans = getStreamStartIdx(str, 14)
        log.Println("Part2: ", ans)
    }
}

func getStreamStartIdx(str string, packetLen int) int{
    start := 0

    packetStartSeq := make([]string, 0)
    curSeq := ""
    for start = 0; start < len(str); start++ {
        end := packetLen

        // take (up to) 4 byte slice
        if start + end > len(str) {
            end = len(str) - start
            end = start + end
        } else {
            end = start + end
        }

        curSeq = str[start:end]


        for _,c := range curSeq {
            if !existsInSlice(string(c), packetStartSeq) {
                packetStartSeq = append(packetStartSeq, string(c))
                if len(packetStartSeq) == packetLen {
                    return start + packetLen
                }
            }
        }
        packetStartSeq = make([]string, 0)
    }

    return -1
}

func existsInSlice(needle string, haystack []string) bool {
    for _, h := range haystack {
        if h == needle {
            return true
        }
    }
    return false
}

func getInputFileLines(fileName string) []string {
    streams := make([]string,0)
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
