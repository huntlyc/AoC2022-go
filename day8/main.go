package main

import (
    "fmt"
	"bufio"
	"log"
	"os"
	"strconv"
)

type dir struct {
	name string
	size int
}

func main() {
	//grid := getInputFileLines("input-test.txt")
	grid := getInputFileLines("input.txt")
    visCount := 0


    seen := make([]string, 0)
    scenicScore := 0

    for rIdx, row := range grid {
        for cIdx, cell := range row {
            //create up down arrays.  This solution is 10x N... LOOP ALL THE ThIngS... Push the CPU - TO THE MAX
            topToBottom := make([]int, 0)
            for _,tmpRow := range grid{
                topToBottom = append(topToBottom, tmpRow[cIdx])
            }

            left := row[0:cIdx]
            right := row[cIdx+1:]
            top := topToBottom[0:rIdx]
            bottom := topToBottom[rIdx+1:]

            newScenicScores := calcScenicScore(cell, rIdx, cIdx, left, right, top, bottom)
            if newScenicScores > scenicScore {
                scenicScore = newScenicScores
            }

            // check left/right
            if isHighest(cell, left) || isHighest(cell, right) || isHighest(cell, top) || isHighest(cell, bottom){
                if !isSeen(seen, cIdx, rIdx) {
                    visCount++
                    seen = append(seen, fmt.Sprintf("%d%d", cIdx, rIdx))
                }
            }
        }
    }

    log.Println("Pt1: ", visCount)
    log.Println("Pt2: ", scenicScore)
}

func calcScenicScore(curH int, r int, c int, left []int, right []int, top []int, bottom []int) int{
    ls, rs, ts, bs := 0, 0, 0, 0


    if len(top) > 0 {
        for i := len(top) -1; i >= 0; i-- {
            if top[i] < curH {
                ts++
            }else if top[i] >= curH{
                ts++
                break
            }
        }
    }else {
        return 0
    }


    if len(left) > 0 {
        for i := len(left) -1; i >= 0; i-- {
            if left[i] < curH {
                ls++
            }else if left[i] >= curH{
                ls++
                break
            }
        }
    }else {
        return 0
    }

    if len(bottom) > 0 {
        for i := 0; i < len(bottom); i++ {
            if bottom[i] < curH {
                bs++
            }else if bottom[i] >= curH{
                bs++
                break
            }
        }
    }else {
        return 0
    }

    if len(right) > 0 {
        for i := 0; i < len(right); i++ {
            if right[i] < curH {
                rs++
            }else if right[i] >= curH{
                rs++
                break
            }
        }
    }else{
        return 0
    }

    if ts == 0 {
        ts = 1
    }

    if ls == 0 {
        ls = 1
    }

    if rs == 0 {
        rs = 1
    }

    if bs == 0 {
        bs = 1
    }

    return ts * ls * rs * bs
}

func isSeen(seen []string, cIdx, rIdx int) bool {
    for _,tmp := range seen {
        val := fmt.Sprintf("%d,%d", cIdx, rIdx)
        if tmp == val{
            return true
        }
    }
    return false
}

func isHighest(needle int, haystack []int) bool {
    if len(haystack) > 0 {
        for _,v := range haystack {
            if v >= needle {
                return false
            }
        }
    }

    return true
}

func getInputFileLines(fileName string) [][]int {
	treeRows := make([][]int, 0)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
        tmp := make([]int, 0)
        for _,t := range sc.Text() {
            v,_ := strconv.Atoi(string(t))
            tmp = append(tmp, v)
        }
        treeRows = append(treeRows, tmp)
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return treeRows
}
