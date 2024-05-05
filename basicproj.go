package main

import (
	"fmt"
    "os"
)


func main(){
    var n int
    fmt.Scanf("%d",&n)

    i := 1
    for i <= n{
        fizzbuzz(i)
        i++
    }

    os.Exit(0)
}

func fizzbuzz(i int) {
    if i % 5 == 0 && i % 3 == 0 {
        fmt.Print(" FizzBuzz")
    } else if i % 3 == 0 {
        fmt.Print(" Fizz")
    } else if i % 5 == 0 {
        fmt.Print(" Buzz")
    } else {
    fmt.Print(" ",i)
    }
}

func isMT[T any](sl []T) bool{
    return len(sl)==0
}

func sqrt[T int | float32 | float64](sl []T) T{
    var sq T
    return sq
}
