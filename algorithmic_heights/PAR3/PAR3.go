package main

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
  "strings"
)


func LoadData(file_path string) (int, []int) {

  f, err := os.Open(file_path)
  if err != nil {
    panic(err)
  }

  content := [][]int{}
  scanner := bufio.NewScanner(f)
  buf := make([]byte, 0, 64*1024)
  scanner.Buffer(buf, 1024*1024)
  for scanner.Scan() {
    int_strs := strings.Fields(scanner.Text())
    ints := []int{}
    for _, str := range int_strs {
      i, _ := strconv.Atoi(str)
      ints = append(ints, i)
    }
    content = append(content, ints)
  }

  n := content[0][0]
  array := content[1]
  return n, array
}


func Partition(A []int) ([]int) {

  q := A[0]
  low := []int{}
  high := []int{}
  equal := []int{q}

  for _, a := range A[1:len(A)] {
    if a < q {
      low = append(low, a)
    } else if a > q {
      high = append(high, a)
    } else {
      equal = append(equal, a)
    }
  }

  par := append(low, equal...)
  par = append(par, high...)
  return par
}


func IntSliceToString(ints []int) (string) {

  int_strs := []string{}
  for _, p := range ints {
    int_strs = append(int_strs, strconv.Itoa(p))
  }

  str := strings.Join(int_strs, " ")
  return str
}


func Write2File(file_path string, content []string) {
  f, err := os.Create(file_path)
  if err != nil {
      panic(err)
  }
  defer f.Close()

  for _, c := range content {
     fmt.Fprintln(f, c)
  }
}


func main () {

  _, A := LoadData("rosalind_par3.txt")
  par := Partition(A)
  Write2File("par_output.txt", []string{IntSliceToString(par)})
}