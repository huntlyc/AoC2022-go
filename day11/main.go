package main

import (
	"bufio"
	"log"
	"os"
    "math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items                    []int
	worryMultiplier          int
	worryOperand             string
	worryTest                int
	nextMonkeyWorryTestTrue  int
	nextMonkeyWorryTestFalse int
	inspectionCount          int
}

func main() {
	monkies := parseMonkiesFromFile("input-test.txt")
	//monkies := parseMonkiesFromFile("input.txt")
	playGamePt1(monkies)
}

func playGamePt1(monkies []Monkey) {
	rounds := 20
	// 1 round is one loop round all monkies
	for ri := 0; ri < rounds; ri++ {
		for mi := 0; mi < len(monkies); mi++ {
			monkey := &monkies[mi]

			// monkey checks each item
			for ii := 0; ii < len(monkey.items); ii++ {
				monkey.inspectionCount += 1

				worryLevel := monkey.items[ii]

				worry := monkey.worryMultiplier
				if worry == -1 {
					worry = worryLevel
				}

				// update the worry level
				if monkey.worryOperand == "*" {
					worryLevel *= worry
				} else if monkey.worryOperand == "+" {
					worryLevel += worry
				}

				// after checking divide by three to the nearest int
				worryLevel = int(math.Floor(float64(worryLevel) / 3)) // pt1

				if worryLevel % monkey.worryTest == 0 {
					monkies[monkey.nextMonkeyWorryTestTrue].items = append(monkies[monkey.nextMonkeyWorryTestTrue].items, worryLevel)
				} else {
					monkies[monkey.nextMonkeyWorryTestFalse].items = append(monkies[monkey.nextMonkeyWorryTestFalse].items, worryLevel)
				}
			}
			// monkey has thrown all items elsewhere
			monkey.items = make([]int, 0)
		}
	}

	inspectionLevels := make([]int, 0)
	for _, m := range monkies {
		inspectionLevels = append(inspectionLevels, m.inspectionCount)
	}

	sort.Ints(inspectionLevels)

	log.Println(inspectionLevels[len(inspectionLevels)-2] * inspectionLevels[len(inspectionLevels)-1])

	//order monkies by inspectionCount
}

func parseMonkiesFromFile(fileName string) []Monkey {
	monkies := make([]Monkey, 0)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	monkey := Monkey{}
	for sc.Scan() {
		line := sc.Text()

		newMonkey := regexp.MustCompile(`Monkey (\d+):`)
		items := regexp.MustCompile(`Starting items: (.*)`)
		worryMultiplier := regexp.MustCompile(`Operation: new = old (.) (.*)`)
		worryTest := regexp.MustCompile(`Test: divisible by (\d+)`)
		nextMonkeyWorryTestTrue := regexp.MustCompile(`If true: throw to monkey (\d+)`)
		nextMonkeyWorryTestFalse := regexp.MustCompile(`If false: throw to monkey (\d+)`)

		if newMonkey.MatchString(line) {
			res := newMonkey.FindAllStringSubmatch(line, -1)
			tmp, _ := strconv.Atoi(res[0][1])

			if tmp > 0 {
				monkies = append(monkies, monkey)
				monkey = Monkey{}
			}
		} else if items.MatchString(line) {
			res := items.FindAllStringSubmatch(line, -1)
			monkeyItems := make([]int, 0)
			tmpStrItems := strings.Split(res[0][1], ", ")
			for _, v := range tmpStrItems {
				tmp, _ := strconv.Atoi(v)
				monkeyItems = append(monkeyItems, tmp)
			}
			monkey.items = monkeyItems
		} else if worryMultiplier.MatchString(line) {

			res := worryMultiplier.FindAllStringSubmatch(line, -1)
			monkey.worryOperand = res[0][1]

			tmp, err := strconv.Atoi(res[0][2])
			if err != nil { // 'old'
				tmp = -1
			}
			monkey.worryMultiplier = tmp
		} else if worryTest.MatchString(line) {
			res := worryTest.FindAllStringSubmatch(line, -1)
			tmp, _ := strconv.Atoi(res[0][1])
			monkey.worryTest = tmp
		} else if nextMonkeyWorryTestTrue.MatchString(line) {
			res := nextMonkeyWorryTestTrue.FindAllStringSubmatch(line, -1)
			tmp, _ := strconv.Atoi(res[0][1])
			monkey.nextMonkeyWorryTestTrue = tmp
		} else if nextMonkeyWorryTestFalse.MatchString(line) {
			res := nextMonkeyWorryTestFalse.FindAllStringSubmatch(line, -1)
			tmp, _ := strconv.Atoi(res[0][1])
			monkey.nextMonkeyWorryTestFalse = tmp
		}
	}
	// add last monkey to list
	monkies = append(monkies, monkey)

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return monkies
}
