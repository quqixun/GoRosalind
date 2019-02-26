package main

import (
    "fmt"
    "bufio"
    "strconv"
    "strings"
    "net/http"
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


func load_protein_ids(text string) ([]string) {

    protein_ids := []string{}

    scanner := bufio.NewScanner(strings.NewReader(text))
    for scanner.Scan() {
        protein_ids = append(protein_ids, scanner.Text())
    }

    return protein_ids
}


func request_uniprot(protein string) (string, bool) {

    url := "https://www.uniprot.org/uniprot/" + protein + ".fasta"
    resp, resp_err := http.Get(url)
    if resp_err != nil {
        return "", false
    }

    info, info_err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if info_err != nil {
        return "", false
    }

    sequence := ""

    info_str := string(info)
    scanner := bufio.NewScanner(strings.NewReader(info_str))
    for scanner.Scan() {
        line_text := scanner.Text()
        if strings.Index(line_text, ">") == -1{
            sequence += line_text            
        }
    }

    return sequence, true
}


func find_positions(sequence string) ([]int) {

    positions := []int{}

    for i := 0; i <= len(sequence) - 4; i++ {
        sub := sequence[i:i + 4]
        if string(sub[0]) == "N" &&
           string(sub[1]) != "P" &&
           (string(sub[2]) == "S" || string(sub[2]) == "T") &&
           string(sub[3]) != "P" {
           positions = append(positions, i + 1) 
        }
    }

    return positions
}


func slice2string(slice []int) (string) {

    slice_temp := []string{}
    for _, item := range slice {
        slice_temp = append(slice_temp, strconv.Itoa(item))
    }

    slice_str := strings.Join(slice_temp, " ")
    return slice_str
}


func main() {

    text, is_load_success := load_text("rosalind_mprt.txt")

    if is_load_success {
        protein_ids := load_protein_ids(text)

        for _, pid := range protein_ids {
            sequence, is_request_success := request_uniprot(pid)

            if is_request_success {                
                positions := find_positions(sequence)
                if len(positions) != 0 {
                    fmt.Println(pid)
                    fmt.Println(slice2string(positions))
                }
            }
        }
    }
}
