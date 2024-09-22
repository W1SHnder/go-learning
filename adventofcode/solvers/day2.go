package solvers

import (
    "fmt" 
    "os"
    "bufio"
    "regexp"
    "strconv"
)

func Day2Part1(file_path string) {
    file, err := os.Open(file_path)
    if err != nil {
        fmt.Print("Encountered error opening file at: ", file_path)
        return 
    }
    //Number of cubes in each bag in the order {red, green, blue}
    var bag = []int{12, 13, 14}

    var regex_string = `^Game ([0-9]+):|([0-9]+) (red|green|blue)` 
    rx := regexp.MustCompile(regex_string) 
    
    scanner := bufio.NewScanner(file)
    var sum = 0 
    var game_id = -1
    var is_valid = true
    for scanner.Scan() {
        line := scanner.Bytes()

        matches := rx.FindAllStringSubmatch(string(line), -1)
        /* 
        for _, match := range matches {
            fmt.Println("Match 1: ", match[2], "Match 2: ", match[3])
        }
        */
        
        for i, match := range matches {
            if i == 0 {
                game_id, err = strconv.Atoi(match[1])
                if err != nil {
                    fmt.Println("Parsing Error: Malformed input")
                    return
                }
            } else {
                switch match[3] {
                case "red":
                    cube_num, err := strconv.Atoi(match[2])
                    if err != nil {
                        fmt.Println("Parsing Error: Malformed input")
                        _ = file.Close()
                        return
                    }
                    if cube_num > bag[0] { is_valid = false }
                    case "green":
                        cube_num, err := strconv.Atoi(match[2])
                        if err != nil {
                            fmt.Println("Parsing Error: Malformed input")
                            _ = file.Close()
                            return
                        }
                        if cube_num > bag[1] { is_valid = false }
                    case "blue":
                        cube_num, err := strconv.Atoi(match[2])
                        if err != nil {
                            fmt.Println("Parsing Error: Malformed input")
                            _ = file.Close()
                            return
                        }
                        if cube_num > bag[2] { is_valid = false }
                }
            } //if
        } //for 
        if is_valid && game_id > -1 {
            sum += game_id 
        }
        is_valid = true
        game_id = -1
    } //for
    _ = file.Close()
    fmt.Println("Sum: ", sum)
} //Day2Part1


func Day2Part2(file_path string) {
    file, err := os.Open(file_path)
    if err != nil {
        fmt.Print("Encountered error opening file at: ", file_path)
        return 
    }
    //Number of cubes in each bag in the order {red, green, blue}
    var bag = []int{0, 0, 0}

    var regex_string = `^Game ([0-9]+):|([0-9]+) (red|green|blue)` 
    rx := regexp.MustCompile(regex_string) 
    
    scanner := bufio.NewScanner(file)
    var sum = 0 
    //var game_id = -1
    for scanner.Scan() {
        line := scanner.Bytes()

        matches := rx.FindAllStringSubmatch(string(line), -1)
        /* 
        for _, match := range matches {
            fmt.Println("Match 1: ", match[2], "Match 2: ", match[3])
        }
        */
        
        for i, match := range matches {
            if i == 0 {
                /*
                game_id, err = strconv.Atoi(match[1])
                if err != nil {
                    fmt.Println("Parsing Error: Malformed input")
                    return
                }
                */
                continue
            } else {
                switch match[3] {
                case "red":
                    cube_num, err := strconv.Atoi(match[2])
                    if err != nil {
                        fmt.Println("Parsing Error: Malformed input")
                        _ = file.Close()
                        return
                    }
                    if cube_num > bag[0] { bag[0] = cube_num }
                    case "green":
                        cube_num, err := strconv.Atoi(match[2])
                        if err != nil {
                            fmt.Println("Parsing Error: Malformed input")
                            _ = file.Close()
                            return
                        }
                        if cube_num > bag[1] { bag[1] = cube_num }
                    case "blue":
                        cube_num, err := strconv.Atoi(match[2])
                        if err != nil {
                            fmt.Println("Parsing Error: Malformed input")
                            _ = file.Close()
                            return
                        }
                        if cube_num > bag[2] { bag[2] = cube_num }
                }
            } //if
        } //for 
        sum +=  bag[0] * bag[1] * bag[2]
        bag[0] = 0
        bag[1] = 0
        bag[2] = 0
    } //for
    _ = file.Close()
    fmt.Println("Sum: ", sum)
} //Day2Part1
