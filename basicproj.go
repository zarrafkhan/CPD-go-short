package main

import (
	"fmt"
    "os"
    // tea "github.com/charmbracelet/bubbletea"
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
