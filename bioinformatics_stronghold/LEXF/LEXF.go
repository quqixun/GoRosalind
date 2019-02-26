package main

import (
    "fmt"
    "bufio"
    "strconv"
    "strings"
    "io/ioutil"
)


func load_text(file_path string) (string, bool) {

    b, err := ioutil.ReadFile(file_path)
    if err != nil {
        fmt.Println(err)
        return "", false
    }

    text := string(b)
    return text, true
}


func read_lines(text string) ([]string, int) {

    lines := []string{}
    scanner := bufio.NewScanner(strings.NewReader(text))
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    symbols := strings.Split(lines[0], " ")
    n, _ := strconv.Atoi(lines[1])

    return symbols, n
}


func NextIndex(ix []int, lens int) {
    for j := len(ix) - 1; j >= 0; j-- {
        ix[j]++
        if j == 0 || ix[j] < lens {
            return
        }
        ix[j] = 0
    }
}


func cartesian_product(symbols []string, n int) ([]string) {

    result := []string{}

    r := ""
    lens := len(symbols)
    for ix := make([]int, n); ix[0] < lens; NextIndex(ix, lens) {
        for _, j := range ix {
            r += symbols[j]
        }
        result = append(result, r)
        r = ""
    }

    return result
}


func main() {
    
    text, is_load_success := load_text("rosalind_lexf.txt")

    if is_load_success {
        symbols, n := read_lines(text)
        result := cartesian_product(symbols, n)
        for _, r := range result {
            fmt.Println(r)
        }
    }
}