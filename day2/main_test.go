package main

import (
    "testing"
)

func TestGamesFromInputPart1(t *testing.T){
    actual := GamesFromInputPart1([]string{"A Y","B X","C Z"})
    expexted := []Game1{{"R","P"},{"P","R"},{"S","S"}}

    for i,v := range expexted {
        if actual[i] != v {
            t.Fatalf("expected %s, got %s", v, actual[i])
        }
    }
}

func TestCalculatePointsPart1(t *testing.T){
    games := []Game1{{"R","P"},{"P","R"},{"S","S"}}
    actual := CalculatePointsPart1(games)
    expexted := 15

    if actual != expexted {
        t.Fatalf("expected %d, got %d", expexted, actual)
    }
}

func TestGamesFromInputPart2(t *testing.T){
    actual := GamesFromInputPart2([]string{"A Y","B X","C Z"})
    expexted := []Game2{{"R","D"},{"P","L"},{"S","W"}}

    for i,v := range expexted {
        if actual[i] != v {
            t.Fatalf("expected %s, got %s", v, actual[i])
        }
    }
}

func TestCalculatePointsPart2(t *testing.T){
    games := []Game2{{"R","D"},{"P","L"},{"S","W"}}
    actual := CalculatePointsPart2(games)
    expexted := 12

    if actual != expexted {
        t.Fatalf("expected %d, got %d", expexted, actual)
    }
}
