package main

import (
    "fmt"
    "strings"
    "strconv"
    "io/ioutil"
)


func read_n_and_k(file_path string) (int, int) {

    b, err := ioutil.ReadFile(file_path)
    if err != nil {
        fmt.Println(err)
        return -1, -1
    }

    text := string(b)
    numbers := strings.Fields(text)

    n, err_n := strconv.Atoi(numbers[0])
    k, err_k := strconv.Atoi(numbers[1])

    if err_n != nil || err_k != nil {
        fmt.Println(err_n, err_k)
        return -1, -1
    }

    return n, k
}


func fibonacci_rabbits(n int, k int) (int) {

    if n <= 0 {
        return 0
    } else if n <= 2 {
        return 1
    } else {
        fr := [2]int{1, 1}
        for i := 3; i <= n; i++ {
            temp := fr[0]
            fr[0] = fr[1]
            fr[1] = temp * k + fr[1]
        }
        return fr[1]
    }
}


func main() {

    n, k := read_n_and_k("rosalind_fib.txt")

    if n != -1 && k != -1 {
        n_rabbits := fibonacci_rabbits(n, k)
        fmt.Println(n_rabbits)
    }
}