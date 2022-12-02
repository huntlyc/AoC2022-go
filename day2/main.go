package main

import (
    "os"
    "bufio"
    "log"
    "strings"
)

const (
    POINTS_ROCK = 1
    POINTS_PAPER = 2
    POINTS_SCISSORS = 3

    POINTS_WIN = 6
    POINTS_DRAW = 3
    POINTS_LOSE = 0
)


type Game1 struct{
    opponentMove string
    playerMove string
}

type Game2 struct{
    opponentMove string
    playerNeedsTo string
}


func main(){
    //inputLines := getInputFileLines("input-test.txt")
    inputLines := getInputFileLines("input.txt")
    games := GamesFromInputPart1(inputLines)
    points := CalculatePointsPart1(games)
    log.Println("Part1: ", points)

    games2 := GamesFromInputPart2(inputLines)
    points2 := CalculatePointsPart2(games2)
    log.Println("Part2: ", points2)
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

func GamesFromInputPart1(lines []string) []Game1 {
    games := make([]Game1, 0)
    for _, game := range lines {

        opponentMove := ""
        playerMove := ""

        // normalise data into R, P, or S
        opponentAndPlayerMoves := strings.Split(game, " ")
        switch opponentAndPlayerMoves[0]{
        case "A": opponentMove = "R"
        case "B": opponentMove = "P"
        case "C": opponentMove = "S"
        }
        switch opponentAndPlayerMoves[1]{
        case "X": playerMove = "R"
        case "Y": playerMove = "P"
        case "Z": playerMove = "S"
        }

        games = append(games, Game1{opponentMove, playerMove})
    }

    return games
}


func CalculatePointsPart1(gameList []Game1) int {
    playerPoints := 0

    for _, game := range gameList {
        switch game.opponentMove{
        case "R":
            switch game.playerMove{
            case "R":
                playerPoints += POINTS_DRAW
                playerPoints += POINTS_ROCK
            case "P":
                playerPoints += POINTS_WIN
                playerPoints += POINTS_PAPER
            case "S":
                playerPoints += POINTS_LOSE
                playerPoints += POINTS_SCISSORS
            }
        case "P":
            switch game.playerMove{
            case "P":
                playerPoints += POINTS_DRAW
                playerPoints += POINTS_PAPER
            case "S":
                playerPoints += POINTS_WIN
                playerPoints += POINTS_SCISSORS
            case "R":
                playerPoints += POINTS_LOSE
                playerPoints += POINTS_ROCK
            }
        case "S":
            switch game.playerMove{
            case "S":
                playerPoints += POINTS_DRAW
                playerPoints += POINTS_SCISSORS
            case "R":
                playerPoints += POINTS_WIN
                playerPoints += POINTS_ROCK
            case "P":
                playerPoints += POINTS_LOSE
                playerPoints += POINTS_PAPER
            }
        }
    }

    return playerPoints
}

func GamesFromInputPart2(lines []string) []Game2 {
    games := make([]Game2, 0)
    for _, game := range lines {

        opponentMove := ""
        playerNeedsTo := ""

        opponentAndPlayerMoves := strings.Split(game, " ")
        switch opponentAndPlayerMoves[0]{
        case "A": opponentMove = "R"
        case "B": opponentMove = "P"
        case "C": opponentMove = "S"
        }
        switch opponentAndPlayerMoves[1]{
        case "X": playerNeedsTo = "L"
        case "Y": playerNeedsTo = "D"
        case "Z": playerNeedsTo = "W"
        }

        games = append(games, Game2{opponentMove, playerNeedsTo})
    }

    return games
}


func CalculatePointsPart2(gameList []Game2) int {
    playerPoints := 0

    for _, game := range gameList {
        switch game.opponentMove{
        case "R":
            switch game.playerNeedsTo{
            case "D":
                playerPoints += POINTS_DRAW
                playerPoints += POINTS_ROCK
            case "W":
                playerPoints += POINTS_WIN
                playerPoints += POINTS_PAPER
            case "L":
                playerPoints += POINTS_LOSE
                playerPoints += POINTS_SCISSORS
            }
        case "P":
            switch game.playerNeedsTo{
            case "D":
                playerPoints += POINTS_DRAW
                playerPoints += POINTS_PAPER
            case "W":
                playerPoints += POINTS_WIN
                playerPoints += POINTS_SCISSORS
            case "L":
                playerPoints += POINTS_LOSE
                playerPoints += POINTS_ROCK
            }
        case "S":
            switch game.playerNeedsTo{
            case "D":
                playerPoints += POINTS_DRAW
                playerPoints += POINTS_SCISSORS
            case "W":
                playerPoints += POINTS_WIN
                playerPoints += POINTS_ROCK
            case "L":
                playerPoints += POINTS_LOSE
                playerPoints += POINTS_PAPER
            }
        }
    }

    return playerPoints
}

