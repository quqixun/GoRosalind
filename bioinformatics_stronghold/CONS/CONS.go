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


func load_samples(text string) (map[string]string) {

    samples := make(map[string]string)

    id := ""
    scanner := bufio.NewScanner(strings.NewReader(text))
    for scanner.Scan() {
        line_text := scanner.Text()
        i := strings.Index(line_text, ">")
        if i > -1 {
            id = line_text[1:]
            samples[id] = ""
        } else {
            if id != "" {
                samples[id] += line_text
            }
        }
    }

    return samples
}


func get_profile(DNAs_map map[string]string) (map[string][]int) {

    DNAs_slice := []string{}
    for _, v := range DNAs_map {
        DNAs_slice = append(DNAs_slice, v)
    }


    A := []int{}
    C := []int{}
    G := []int{}
    T := []int{}

    n_cols := len(DNAs_slice[0])
    for i := 0; i < n_cols; i++ {
        col := ""
        for _, DNA := range DNAs_slice {
            col += string(DNA[i])
        }

        A = append(A, strings.Count(col, "A"))
        C = append(C, strings.Count(col, "C"))
        G = append(G, strings.Count(col, "G"))
        T = append(T, strings.Count(col, "T"))
    }

    profile := make(map[string][]int)
    profile["A"] = A
    profile["C"] = C
    profile["G"] = G
    profile["T"] = T
    return profile
}


func get_consensus(profile map[string][]int) (string) {

    consensus := ""
    consensus_len := len(profile["A"])
    keys := []string{"A", "C", "G", "T"}

    for i := 0; i < consensus_len; i++ {

        max_num, max_key := 0, ""
        for _, k := range keys {
            if profile[k][i] > max_num {
                max_num = profile[k][i]
                max_key = k
            }
        }
        consensus += max_key
    }

    return consensus
}


func slice2string(slices []int) (string) {

    slice_temp := []string{}
    for _, slice := range slices {
        slice_temp = append(slice_temp, strconv.Itoa(slice))
    }

    slices_str := strings.Join(slice_temp, " ")
    return slices_str
}


func main() {

    text, is_text_success := load_text("rosalind_cons.txt")
    if is_text_success {
        DNAs := load_samples(text)
        profile := get_profile(DNAs)
        consensus := get_consensus(profile)

        fmt.Println(consensus)
        fmt.Println("A: " + slice2string(profile["A"]))
        fmt.Println("C: " + slice2string(profile["C"]))
        fmt.Println("G: " + slice2string(profile["G"]))
        fmt.Println("T: " + slice2string(profile["T"]))
    }
}