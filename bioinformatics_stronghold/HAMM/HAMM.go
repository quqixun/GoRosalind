package main

import (
    "fmt"
    "bufio"
    "strings"
    "io/ioutil"
)


func load_text(file_path string) (string) {

    b, err := ioutil.ReadFile(file_path)
    if err != nil {
        fmt.Println(err)
        return "Error"
    }

    text := string(b)
    return text
}


func load_samples(text string) ([]string) {

    samples := make([]string, 2)

    index := 0
    scanner := bufio.NewScanner(strings.NewReader(text))
    for scanner.Scan() {

        if index == 2 {
            break
        }

        samples[index] = scanner.Text()
        index += 1
    }

    return samples
}


func count_point_mutations(samples []string) (int) {

    sequence0 := samples[0]
    sequence1 := samples[1]

    if len(sequence0) != len(sequence1) {
        return -1
    }

    cpm := 0
    for i, v := range sequence0 {
        if string(v) != string(sequence1[i]) {
            cpm += 1
        }
    }

    return cpm
}


func main() {

    text := load_text("rosalind_hamm.txt")

    if text != "error" {
        samples := load_samples(text)
        cpm := count_point_mutations(samples)
        fmt.Println(cpm)
    }
}