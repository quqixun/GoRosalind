package main

import (
    "fmt"
    "io/ioutil"
)


func read_text(file_path string) (string) {

    b, err := ioutil.ReadFile(file_path)
    if err != nil {
        fmt.Println(err)
        return "Error"
    }

    text := string(b)
    return text
}


func reverse_complement(DNA string) (string) {

    DNA_rc := ""
    for _, s := range DNA {
        switch s {
            case 'A': DNA_rc = "T" + DNA_rc
            case 'T': DNA_rc = "A" + DNA_rc
            case 'C': DNA_rc = "G" + DNA_rc
            case 'G': DNA_rc = "C" + DNA_rc
        }
    }
    return DNA_rc
}


func main() {

    DNA := read_text("rosalind_revc.txt")

    if DNA != "error" {
        DNA_rc := reverse_complement(DNA)
        fmt.Println(DNA_rc)
    }
}