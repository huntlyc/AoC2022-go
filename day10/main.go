package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	x := 1                  // value of x
	xvals := make([]int, 0) // historical values of x
	xvals = append(xvals, x)
	sumTally := 0

	totalCycles := 0

	 cmds := getInputFileLines("input.txt")
	//cmds := getInputFileLines("input-test.txt")
	//cmds := getInputFileLines("input2")


	for _, rawcmd := range cmds {

        cmd := strings.Split(rawcmd, " ")
        switch cmd[0] {
        case "noop":
            totalCycles += 1
            sumTally = check(totalCycles,xvals,x,sumTally)
        case "addx":
            totalCycles += 1
            //check
            sumTally = check(totalCycles,xvals,x,sumTally)


            // on the second cycle add the x
            totalCycles += 1
            //check
            sumTally = check(totalCycles,xvals,x,sumTally)

			if len(cmd) == 2 {
                newX, _ := strconv.Atoi(cmd[1])
                x += newX
			}
        }
	}
	log.Println(sumTally)
}

func check(totalCycles int, xvals []int, x int, sumTally int) int{
    if totalCycles  == 20 || ((totalCycles-20)%40) == 0 {
        tmpx := 0
        for _, v := range xvals {
            tmpx += v
        }

        sigStren := x * totalCycles
        sumTally += sigStren
        log.Println(x, totalCycles, tmpx, sumTally)
        return sumTally
    }
    return sumTally
}



func getInputFileLines(fileName string) []string {
	cmds := make([]string, 0)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		cmds = append(cmds, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return cmds
}
