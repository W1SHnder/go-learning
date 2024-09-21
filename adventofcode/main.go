package main

import (
    "adventofcode/algos"
)

func main() {
    /* 
    if len(os.Args) != 2 {
        fmt.Println("Useage: ./main <input>")
        return
    }
    */
    
    //var strs = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    
    var strs = []string{"a", "ab", "bc", "bca", "c", "caa"}

    var tree = algos.NewAhoTree(strs) 
    algos.AhoUglyPrint(tree)
} //main
