package main
import (
    "testing"
)
func TestHighest(t *testing.T){
    actual := Highest([]int{3,2,1})
    expexted := 3
    if actual != expexted {
        t.Fatalf("expected %d, got %d", expexted, actual)
    }
}

func TestSumOfTopThree(t *testing.T){
    actual := SumOfTopThree([]int{3,2,1})
    expexted := 6
    if actual != expexted {
        t.Fatalf("expected %d, got %d", expexted, actual)
    }
}
