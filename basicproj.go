package main

import "fmt"
import "unicode/utf8"

func main(){
    //date types
    //int + float just like c frfr 
    var declare uint32 = 34 //var name type

    fmt.Println(declare)
    
    str  := "frfr"
    fmt.Println(str)

    fmt.Println(len(str))

    fmt.Println(utf8.RuneCountInString(str))

    inf := "intferred type with short form"
    fmt.Println(inf)

    new(str)

    arr := [4]uint16{0,2,4,6}

    fmt.Println(arr[2])
    fmt.Println(arr[0:2])

    newMap := map[string]uint32{"test":32,"news":42}

    fmt.Println(newMap["news"])
    first, second := newMap["rand"]
    fmt.Println(first, second)
//  maps will still return smth even if the key does not exist


}

func new(in string){
    fmt.Println(in)
}
