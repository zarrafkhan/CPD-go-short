package main

import "fmt"
import "unicode/utf8"

func main(){
    //date types
    //int + float just like c frfr 
    var declare uint32 = 34 //var name type

    fmt.Println(declare)
    
    var str string = "frfr"
    fmt.Println(str)


    fmt.Println(len(str))

    fmt.Println(utf8.RuneCountInString(str))

    // var arr [3]int = [3]int{1, 2, 3}
    // var sl []int = []int{1, 2, 3}
    // var st struct { x int; y int } = struct { x int; y int }{1, 2}
    //
    // var p *int = &i
    // var ch chan int = make(chan int)
    // var m map[string]int = make(map[string]int)
    // var iface interface{} = "text" 
    //

    inf := "intferred type with short form"
    fmt.Println(inf)

}
