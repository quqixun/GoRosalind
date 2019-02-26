package main

import (
    "fmt"
    "math"
    "bufio"
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


func load_table(file_path string) (map[string]int, bool) {

    b, err := ioutil.ReadFile(file_path)
    if err != nil {
        fmt.Println(err)
        return nil, false
    }

    text := string(b)
    scanner := bufio.NewScanner(strings.NewReader(text))

    table := make(map[string]int)
    for scanner.Scan() {
        line_text := scanner.Text()
        line_words := strings.Fields(line_text)
        table[line_words[1]] += 1
    }

    return table, true
}


func compute_remainder(protein string, RNA_codon map[string]int) (float64) {

    var multiply_res float64
    multiply_res = 1.0

    for _, c := range protein {
        multiply_res *= float64(RNA_codon[string(c)])
        multiply_res = math.Mod(multiply_res, 1000000.0)
    }

    multiply_res *= float64(RNA_codon["Stop"])
    mod_res := math.Mod(multiply_res, 1000000.0)
    return mod_res
}


func main() {

    protein_text, is_load_success := load_text("rosalind_mrna.txt")
    RNA_codon_table, is_table_sucess := load_table("RNA_codon_table.txt")

    if is_load_success && is_table_sucess {
        mod_res := compute_remainder(protein_text, RNA_codon_table)
        fmt.Println(mod_res)
    }
}
