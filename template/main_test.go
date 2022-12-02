package main

import (
    "testing"
)

func TestAdd(t *testing.T){
    expected := 2
    actual := Add(1,1)
    if actual != expected {
        t.Fatalf("expected %d, got %d", expected, actual)
    }
}


