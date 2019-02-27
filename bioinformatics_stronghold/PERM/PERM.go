package main

import (
  "fmt"
  "strconv"
  "strings"
  "io/ioutil"
)


func read_number(file_path string) (int, bool) {

  b, err_b := ioutil.ReadFile(file_path)
  if err_b != nil {
    fmt.Println(err_b)
    return 0, false
  }

  number, err_n := strconv.Atoi(string(b))
  if err_n != nil {
    fmt.Println(err_n)
    return 0, false
  }

  return number, true
}


func factorial(n int) (r int) {

  if n > 0 {
    r = n * factorial(n - 1)
    return r
  }

  return 1
}


func permutations(n int) (int, [][]int) {

  pn := factorial(n)
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
  return pn, ps
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

  n, is_read_success := read_number("rosalind_perm.txt")    

  if is_read_success {
    pn, ps := permutations(n)
    fmt.Println(pn)
    for _, p := range ps {
      fmt.Println(slice2string(p))
    }
  }
}