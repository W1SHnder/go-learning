package main

import (
    "adventofcode/solvers"
    "fmt"
    "os"
)

func main() {
    
    if len(os.Args) != 2 {
        fmt.Println("Useage: ./main <input>")
        return
    }

    solvers.Day2Part2(os.Args[1])
} //main
