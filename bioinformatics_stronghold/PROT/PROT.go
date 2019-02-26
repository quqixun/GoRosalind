package main

import (
    "fmt"
    "bufio"
    "strings"
    "io/ioutil"
)


func load_table(file_path string) (map[string]string, bool) {

    b, err := ioutil.ReadFile(file_path)
    if err != nil {
        fmt.Println(err)
        return nil, false
    }

    text := string(b)
    scanner := bufio.NewScanner(strings.NewReader(text))

    table := make(map[string]string)
    for scanner.Scan() {
        line_text := scanner.Text()
        line_words := strings.Fields(line_text)
        table[line_words[0]] = line_words[1]
    }

    return table, true
}


func load_text(file_path string) (string, bool) {

    b, err := ioutil.ReadFile(file_path)
    if err != nil {
        fmt.Println(err)
        return "", false
    }

    text := string(b)
    return text, true
}


func translate(sequence string, table map[string]string) (string, bool) {

    ul := 3 // unit length
    protein_sequence := ""
    iter_num := (len(sequence) - len(sequence) % ul) / ul

    for i := 0; i < iter_num; i++ {
        protein := table[sequence[ul * i:ul * (i + 1)]]
        if protein == "" {
            return "", false
        } else if protein == "Stop" {
            break
        } else {
            protein_sequence += protein
        }
    }

    return protein_sequence, true
}


func main() {

    RNA_codon_table, is_table_success := load_table("RNA_codon_table.txt")
    RNA_sequence, is_text_success := load_text("rosalind_prot.txt")

    if is_table_success && is_text_success {
        protein, is_success := translate(RNA_sequence, RNA_codon_table)
        if is_success {
            fmt.Println(protein)
        }
    }
}