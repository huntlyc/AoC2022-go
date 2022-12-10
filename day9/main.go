package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Vec struct {
	x int
	y int
}

type Move struct {
	dir Vec
	num int
}

type Pos struct {
	x int
	y int
}

func main() {
	//moveSet := getInputFileLines("input-test.txt")
	moveSet := getInputFileLines("input.txt")
	//moveSet := getInputFileLines("input2")

	//part 1
	//runSimulation(moveSet, 2)

	//part 1
	runSimulation(moveSet, 10)

}

func runSimulation(moveSet []Move, ropeLength int) {
	rope := make([]Pos, 0)

    tailUniqueVisits := make([]Pos, 0)
    tailUniqueVisits = append(tailUniqueVisits, Pos{0,0})


	for i := 0; i < ropeLength; i++ {
		rope = append(rope, Pos{x: 0, y: 0})
	}

	for _, m := range moveSet {
		for i := 0; i < m.num; i++ {
			rope := moveRope(rope, m.dir)

            tailTip := rope[len(rope)-1]
            if !isVisited(tailUniqueVisits, tailTip) {
                tailUniqueVisits = append(tailUniqueVisits, tailTip)
            }
		}
	}

    log.Println("ans: ",  len(tailUniqueVisits))
}

/**
 * borrowed from rust, not mut so all good ;)
 **/
func signum(n int) int {
    if n > 0 {
        return 1
    }
    if n < 0 {
        return -1
    }
    return 0
}


// for such a simple fn, it took 6 rewrites....
func moveRope(rope []Pos, dir Vec) []Pos {

	for i := 0; i < len(rope); i++ {
		knot := &rope[i]
		if i == 0 {
			knot.x += dir.x
			knot.y += dir.y
		} else {
			head := &rope[i-1]

			xDiff := head.x - knot.x
			yDiff := head.y - knot.y

            absXDiff := math.Abs(float64(xDiff))
            absYDiff := math.Abs(float64(yDiff))


            if absXDiff == 2 || absYDiff == 2 {
                knot.x += signum(xDiff)
                knot.y += signum(yDiff)
            }
		}
	}
	return rope
}

func isTouching(follower Pos, leader Pos) bool {

    xDiff := math.Abs(float64(follower.x - leader.x))
    yDiff := math.Abs(float64(follower.y - leader.y))

    if xDiff > 1 || yDiff > 1 {
        return false
    }

    return true
}

func isVisited(visited []Pos, cur Pos) bool {
	for _, p := range visited {
		if p.x == cur.x && p.y == cur.y {
			return true
		}
	}
	return false
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
		curMove := Move{dir: translateDirToVec(m[0]), num: v}
		moveSet = append(moveSet, curMove)
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return moveSet
}

func translateDirToVec(dir string) Vec {
	switch dir {
	case "R":
		return Vec{x: 1, y: 0}
	case "L":
		return Vec{x: -1, y: 0}
	case "U":
		return Vec{x: 0, y: 1}
	case "D":
		return Vec{x: 0, y: -1}
	}

	return Vec{0, 0}
}

func translateVecToDir(vec Vec) string {
    if vec.x == 1  && vec.y == 0{
        return "R"
    }
    if vec.x == -1 && vec.y == 0{
        return "L"
    }
    if vec.x == 0  && vec.y == 1{
        return "U"
    }
    if vec.x == 0  && vec.y == -1{
        return "D"
    }
    return "?"
}
