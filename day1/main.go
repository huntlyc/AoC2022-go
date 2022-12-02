package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main(){
    fmt.Println("Hello")
    file, err := os.Open("./input.txt");
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()


    var totals []int
    cur := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        line := scanner.Text()
        if line == "" {
            totals = append(totals, cur)
            cur = 0
        }else{
            next, err :=  strconv.Atoi(line)
            if err != nil {
                log.Fatal(err)
            }
            cur += next
        }
    }

    sort.Slice(totals, func(i, j int) bool {
        return totals[i] > totals[j]
    })

    fmt.Println("Pt1: ", Highest(totals)) // 70698
    fmt.Println("Pt2: ", SumOfTopThree(totals)) // 206643

}

func Highest(calorieTotals []int) int {
    return calorieTotals[0]
}

func SumOfTopThree(calorieTotals []int ) int{
    sum := 0
    for _,num := range calorieTotals[0:3]{
        sum += num
    }
    return sum
}
