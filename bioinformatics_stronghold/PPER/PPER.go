package main

import (
    "fmt"
    "strconv"
    "strings"
    "math/big"
    "io/ioutil"
)


func load_n_and_k(file_path string) (int, int, bool) {

    n, k := 0, 0

    b, err := ioutil.ReadFile(file_path)
    if err != nil {
        fmt.Println(err)
        return n, k, false
    }

    text := string(b)
    words := strings.Fields(text)
    n, _ = strconv.Atoi(words[0])
    k, _ = strconv.Atoi(words[1])

    return n, k, true
}


func partial_permutations(n, k int) (*big.Int) {

    res := big.NewInt(1)

    for i := n; i >= n - k + 1; i-- {
        res = res.Mul(res, big.NewInt(int64(i)))
    }

    return res
}


func main() {

    n, k, is_load_sucess := load_n_and_k("rosalind_pper.txt")

    if is_load_sucess {
        res := partial_permutations(n, k)
        res_mod := res.Mod(res, big.NewInt(1000000))
        fmt.Println(res_mod.String())
    }

}
