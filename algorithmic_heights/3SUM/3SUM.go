package main

import (
  "fmt"
  "sort"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (int, int, [][]int) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  text := string(b)
  content := [][]int{}
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    int_strs := strings.Fields(scanner.Text())
    ints := []int{}
    for _, str := range int_strs {
      i, _ := strconv.Atoi(str)
      ints = append(ints, i)
    }
    content = append(content, ints)
  }

  k := content[0][0]
  n := content[0][1]
  arrays := content[1:len(content)]
  return k, n, arrays
}


func SliceIndex(t int, ints []int) (int) {

  for i, e := range ints {
    if e == t {
      return i
    }
  }

  return -1
}


func DifferentIndices(arr []int) ([]int) {

  arr_c := make([]int, len(arr))
  copy(arr_c, arr)
  sort.Ints(arr_c)

  indices := make([]int, 3)

  arr_len := len(arr)
  for i := 0; i < arr_len - 2; i++ {
    a := arr_c[i]
    j := i + 1
    k := arr_len - 1
    for {
      if j >= k {
        break
      } else {
        b := arr_c[j]
        c := arr_c[k]
        if a + b + c == 0 {
          indices = []int{
            SliceIndex(a, arr) + 1,
            SliceIndex(b, arr) + 1,
            SliceIndex(c, arr) + 1,
          }
          sort.Ints(indices)
          return indices
        } else if a + b + c > 0 {
          k -= 1
        } else {
          j += 1
        }
      }
    }
  }

  return []int{-1}
}


func IntSliceToString(ints []int) (string) {

  int_strs := []string{}
  for _, p := range ints {
    int_strs = append(int_strs, strconv.Itoa(p))
  }

  str := strings.Join(int_strs, " ")
  return str
}


func main() {

  _, _, arrays := LoadData("rosalind_3sum.txt")

  for _, arr := range arrays {
    res := DifferentIndices(arr)
    fmt.Println(IntSliceToString(res))
  }
}
