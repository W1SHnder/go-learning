package solvers

import (
    "fmt"
    "os"
    "bufio"
)

func Day1Part1(file_path string) {
    var input_file, err = os.Open(file_path)
    if (err != nil) {
        fmt.Println("Error reading file: ", err)
    }
    var reader = bufio.NewReader(input_file)
    var is_first = true
    var last_digit byte
    var cur_num = 0
    var sum = 0
    for {
        b, err := reader.ReadByte()
        if (err != nil) {
            break 
        }
        if (b >= '0' && b <= '9') {
            if is_first {
                cur_num += 10 * int((b - 48))
                last_digit = b
                is_first = false
            } else {
                last_digit = b 
            } 
        } else if (b == '\r' || b == '\n'){
            if !is_first {
                cur_num += int(last_digit - 48)
                sum += cur_num
                cur_num = 0
                is_first = true  
            }
        } else {
            continue 
        } //if
    } 
    fmt.Print(sum)
} //Day1Part1

func Day1Part2(file_path string) {
    var input_file, err = os.Open(file_path)
    if err != nil {
        fmt.Println("Error reading file: ", err)
    }
    
    var reader = bufio.NewReader(input_file)
    var is_first = true
    var last_digit byte
    var cur_num = 0
    var sum = 0
    var num_str = make([]byte, 64)
    var i = 0
    for {
        b, err := reader.ReadByte()
        if err != nil { break; }
        if (b >= '0' && b <= '9') {
            if is_first {
                cur_num += 10 * int(b - 48)
                last_digit = b
                is_first = false
            } else {
                last_digit = b
            }
        } else if (b == '\r' || b == '\n') {
            if !is_first {
                cur_num += int(last_digit - 48)
                sum += cur_num
                cur_num = 0
                is_first = true
                num_str = num_str[0:]
            }
        } else {
            num_str[i] = b
            if len(num_str) > 2 {
                
            }
        } //if
    } //for
} 

func checkSubstrings() {

}
