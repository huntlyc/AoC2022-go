package main

import (
    "fmt"
	"bufio"
	"log"
	"os"
    "regexp"
    "strconv"
    "strings"
)

type elfPair struct {
	firstElf  string
	secondElf string
}

var stacks = make([][]string, 0)

func main() {
    part2();
}

func part1(){
	//file, err := os.Open("input-test.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		line := sc.Text()

        // if we're on a stack line, parse all stack items
        re := regexp.MustCompile(`\[[A-Z]\]`)
        if re.MatchString(line) {
            stackNum := 1
            start := 0
            chunk := 4

            for ((start + chunk) <= len(line)){
                entry := strings.Trim(line[start:(start + chunk)], " []")

                if len(stacks) < stackNum {
                    stacks = append(stacks, make([]string, 0))
                }


                if entry != "" {
                    stacks[stackNum-1] = append(stacks[stackNum-1], entry)
                }

                stackNum += 1


                start += chunk
                if start + chunk >= len(line){
                    chunk -= 1
                }
            }
            stackNum = 1
            start = 0
            chunk = 4
        }



        // check for move command
        re = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
        if re.MatchString(line) {
        printStacks()
            fmt.Println(line)
            res := re.FindAllStringSubmatch(line, -1)
                amount,_ := strconv.Atoi(res[0][1])
                from,_ := strconv.Atoi(res[0][2])
                to,_ := strconv.Atoi(res[0][3])

                from -= 1
                to -= 1

                for i := 0; i < amount; i++ {
                log.Print(amount)
                    cargo := []string{}
                    cargo = append(cargo, stacks[from][0:1]...)

                    newFrom := []string{}
                    newFrom = append(newFrom, stacks[from][1:]...)
                    stacks[from] = newFrom



                    cargo = append(cargo,
                    stacks[to][0:]...)
                    stacks[to] = cargo
                }

            printStacks()
        }
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}



}

func part2(){
	//file, err := os.Open("input-test.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		line := sc.Text()

        // if we're on a stack line, parse all stack items
        re := regexp.MustCompile(`\[[A-Z]\]`)
        if re.MatchString(line) {
            stackNum := 1
            start := 0
            chunk := 4

            for ((start + chunk) <= len(line)){
                entry := strings.Trim(line[start:(start + chunk)], " []")

                if len(stacks) < stackNum {
                    stacks = append(stacks, make([]string, 0))
                }


                if entry != "" {
                    stacks[stackNum-1] = append(stacks[stackNum-1], entry)
                }

                stackNum += 1


                start += chunk
                if start + chunk >= len(line){
                    chunk -= 1
                }
            }
            stackNum = 1
            start = 0
            chunk = 4
        }



        // check for move command
        re = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
        if re.MatchString(line) {
        printStacks()
            fmt.Println(line)
            res := re.FindAllStringSubmatch(line, -1)
            amount,_ := strconv.Atoi(res[0][1])
            from,_ := strconv.Atoi(res[0][2])
            to,_ := strconv.Atoi(res[0][3])

            from -= 1
            to -= 1

            cargo := []string{}
            cargo = append(cargo, stacks[from][0:amount]...)

            newFrom := []string{}
            newFrom = append(newFrom, stacks[from][amount:]...)
            stacks[from] = newFrom



            cargo = append(cargo,
            stacks[to][0:]...)
            stacks[to] = cargo

            printStacks()
        }
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}



}
func printStacks(){
    for si,st := range stacks {
        log.Print(si,st)
    }
}
