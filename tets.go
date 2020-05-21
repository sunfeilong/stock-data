package main

import (
    "fmt"
    "regexp"
)

func main() {
    text := "das<u>name</u>dsa"



    matchs := compile.FindAllSubmatch([]byte(text), 1)

    fmt.Println(matchs[0][2])
    fmt.Println(string(matchs[0][2]))

}
