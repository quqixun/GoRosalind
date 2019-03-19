package main

import (
  "fmt"
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


func DifferentIndices(arr []int) ([]int) {

  arr_len := len(arr)
  for p := 0; p < arr_len - 1; p++ {
    for q := p + 1; q < arr_len; q++ {
      if arr[p] == arr[q] * -1 {
        return []int{p + 1, q + 1}
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

  _, _, arrays := LoadData("rosalind_2sum.txt")

  for _, arr := range arrays {
    res := DifferentIndices(arr)
    fmt.Println(IntSliceToString(res))
  }
}
