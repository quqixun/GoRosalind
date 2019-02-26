package main

import (
    "fmt"
    "strings"
    "io/ioutil"
)


func read_text(file_path string) (string) {

    b, err := ioutil.ReadFile(file_path)
    if err != nil {
        fmt.Println(err)
        return "error"
    }

    text := string(b)
    return text
}


func count_DNA(DNA string) (map[string]int) {

    c := make(map[string]int)
    c["A"] = strings.Count(DNA, "A")
    c["C"] = strings.Count(DNA, "C")
    c["G"] = strings.Count(DNA, "G")
    c["T"] = strings.Count(DNA, "T")
    return c
}


func main() {

    DNA := read_text("rosalind_dna.txt")    

    if DNA != "error" {
        count := count_DNA(DNA)
        fmt.Println(count["A"], count["C"], count["G"], count["T"])
    }
}
