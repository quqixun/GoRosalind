package main

import (
  "fmt"
  "strconv"
  "strings"
  "io/ioutil"
)


func load_n(file_path string) (int) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  tb := string(b)
  n, _ := strconv.Atoi(tb)

  return n
}


func Permutations(n int) ([][]int) {

  ps := [][]int{}

  arr := []int{}
  for i := 1; i <= n; i++ {
    arr = append(arr, i)
  }

  var generate func([]int, int)
  generate = func(arr []int, t int) {
    if t == 1 {
      tmp := make([]int, len(arr))
      copy(tmp, arr)
      ps = append(ps, tmp)
    } else {
      for i := 0; i < t; i++ {
        generate(arr, t - 1)
        if t % 2 == 1 {
          tmp := arr[i]
          arr[i] = arr[t - 1]
          arr[t - 1] = tmp
        } else {
          tmp := arr[0]
          arr[0] = arr[t - 1]
          arr[t - 1] = tmp
        }
      }
    }
  }

  generate(arr, len(arr))
  return ps
}


func rCombinations(p int, n []int, c []int, ccc [][][]int) [][][]int {
    if len(n) == 0 || p <= 0 {
        return ccc
    }
    if len(ccc) == 0 {
        ccc = make([][][]int, p)
    }
    p--
    for i := range n {
        cc := make([]int, len(c)+1)
        copy(cc, c)
        cc[len(cc)-1] = n[i]
        ccc[len(cc)-1] = append(ccc[len(cc)-1], cc)
        ccc = rCombinations(p, n[i+1:], cc, ccc)
    }
    return ccc
}


func Combinations(p int, n []int) [][][]int {
    return rCombinations(p, n, nil, nil)
}


func signed(perm [][]int, n int) ([][]int) {

  signed_perm := make([][]int, len(perm))
  copy(signed_perm, perm)

  idx_list := []int{}
  for i := 0; i < n; i++ {
    idx_list = append(idx_list, i)
  }
  idx_comb := Combinations(n, idx_list)

  for _, ict := range idx_comb {
    for _, c := range ict {
      for _, sp := range perm {
        spc := make([]int, len(sp))
        copy(spc, sp)
        for _, i := range c {
          spc[i] *= -1
        }
        signed_perm = append(signed_perm, spc)
      }
    }
  }

  return signed_perm
}


func signed_permutations(n int) ([][]int) {

  numbers := []int{}
  for i := 1; i <= n; i++ {
    numbers = append(numbers, i)
  }

  perm := Permutations(n)
  signed_perm := signed(perm, n)

  return signed_perm
}


func main() {
  
  n := load_n("rosalind_sign.txt")
  sp := signed_permutations(n)

  fmt.Println(len(sp))
  for _, spe := range sp {
    spe_strs := []string{}
    for _, s := range spe {
      spe_strs = append(spe_strs, strconv.Itoa(s))
    }
    spe_str := strings.Join(spe_strs, " ")
    fmt.Println(spe_str)
  }
}