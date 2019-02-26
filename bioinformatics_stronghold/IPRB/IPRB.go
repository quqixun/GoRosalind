package main

import (
    "fmt"
    "strconv"
    "strings"
    "io/ioutil"
)


func load_kmn(file_path string) (int, int, int) {

    b, err := ioutil.ReadFile(file_path)
    if err != nil {
        fmt.Println(err)
    }

    text := string(b)
    numbers := strings.Fields(text)

    k, _ := strconv.Atoi(numbers[0])
    m, _ := strconv.Atoi(numbers[1])
    n, _ := strconv.Atoi(numbers[2])
    return k, m, n
}


func prop_dominant_allele(k int, m int, n int) (float64) {

    t := float64(k + m + n)

    // yy * yy
    hr1 := float64((n * (n - 1))) / (t * (t - 1))

    // Yy * yy or yy * Yy
    hr2 := float64((m * n)) / (t * (t - 1))

    //Yy * Yy
    hr3 := float64((m * (m - 1))) / (4 * t * (t - 1))

    prop := 1 - (hr1 + hr2 + hr3)
    return prop
}


func main() {

    k, m, n := load_kmn("rosalind_iprb.txt")

    prop := prop_dominant_allele(k, m, n)
    fmt.Printf("%.5f", prop)

}
