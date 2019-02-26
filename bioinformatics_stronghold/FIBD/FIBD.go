package main

import (
    "fmt"
    "strings"
    "strconv"
    "io/ioutil"
)


func read_n_and_m(file_path string) (int, int, bool) {

    b, err := ioutil.ReadFile(file_path)
    if err != nil {
        fmt.Println(err)
        return 0, 0, false
    }

    text := string(b)
    numbers := strings.Fields(text)

    n, err_n := strconv.Atoi(numbers[0])
    m, err_m := strconv.Atoi(numbers[1])

    if err_n != nil || err_m != nil {
        fmt.Println(err_n, err_m)
        return 0, 0, false
    }

    return n, m, true
}


func sum(slice []int) (int) {

    r := 0
    for _, s := range slice {
        r += s
    }

    return r
}


func mortal_fibonacci_rabbits(n int, m int) (int) {

    if n <= 2 {
        return 1
    }

    var Gt_1 []int
    var Gt []int

    for i := 0; i < m; i++ {
        Gt_1 = append(Gt_1, 0)
        Gt = append(Gt, 0)
    }

    Gt_1[1] = 1

    for t := 3; t <= n; t++ {
        for i := 0; i < m; i++ {
            if i == 0 {
                Gt[i] = sum(Gt_1) - Gt_1[0]
            } else if i == m - 1 {
                Gt[i] = Gt_1[m - 2]
            } else {
                Gt[i] = Gt_1[i - 1]
            }
        }
        copy(Gt_1, Gt)
    }

    mfr := sum(Gt)
    return mfr
}


func main() {

    n, m, is_read_success := read_n_and_m("rosalind_fibd.txt")

    if is_read_success {
        mfr := mortal_fibonacci_rabbits(n, m)
        fmt.Println(mfr)
    }
}