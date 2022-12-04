package main

import (
	"bufio"
	"log"
	"os"
)

func main(){
    //rucksacks := getInputFileLines("input-test.txt")
    rucksacks := getInputFileLines("input.txt")
    part1(rucksacks)
    part2(rucksacks)
}

func part2(rucksacks []string) {
    total := 0
    count := 0



    for count < len(rucksacks){

        comp1 := uniqItems(rucksacks[count])
        comp2 := uniqItems(rucksacks[count + 1])
        comp3 := uniqItems(rucksacks[count + 2])


        found2, found3 := false, false
        for _,r1 := range comp1 {
            for _,r2 := range comp2 {
                if r1 == r2 {
                    found2 = true;

                    for _,r3 := range comp3 {
                        if r2 == r3 {
                            found3 = true
                        }
                    }
                }
            }

            if found2 && found3 {
                points := convRuneToPriorityValue(r1)
                total += points
            }

            found2, found3 = false, false
        }

        count += 3
    }

    log.Println("Part2 total: ", total)
}

func part1(rucksacks []string) {
    total := 0
    found := false

    for _,raw := range rucksacks {
        // split into 2 compartments
        mid := len(raw) / 2
        tmp1, tmp2 := raw[0:mid], raw[mid:]
        comp1 := uniqItems(tmp1)
        comp2 := tmp2

        found = false
        for _,c1l := range comp1 {
            for _,c2l := range comp2{
                if found == false && c1l == c2l {
                    found = true
                    v := convRuneToPriorityValue(c1l)

                    total += v
                }
            }
            found = false
        }
    }

    log.Println("Part1 total: ", total)
}

func uniqItems(tmp1 string) string{
    found := false
    comp1 := ""
    for _,l := range tmp1{
        for _,tl := range comp1{
            if l == tl {
                found = true
            }
        }
        if found == false {
            comp1 += string(l)
        }
        found = false
    }
    return comp1
}

func convRuneToPriorityValue(r rune) int{
    // a -> z : 1  -> 26
    // A -> Z : 27 -> 52
    rv := int(r)
    if rv >= 97 {
        rv -= 96
    } else {
        rv -= 38
    }
    return rv
}

func getInputFileLines(fileName string) []string {
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    lines := make([]string, 0)

    for sc.Scan() {
        lines = append(lines, sc.Text())
    }

    if err := sc.Err(); err != nil {
        log.Fatal(err)
    }

    return lines
}
