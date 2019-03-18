package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (int, []int, int, []int) {

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

  n := content[0][0]
  A := content[1][0:n]
  m := content[2][0]
  B := content[3][0:m]
  return n, A, m, B
}


func MakeCopy(arr []int) ([]int) {

  arr_copy := make([]int, len(arr))
  copy(arr_copy, arr)

  return arr_copy
}


func MergeSortedArrays(A, B []int) ([]int) {

  Ac := MakeCopy(A)
  Bc := MakeCopy(B)

  msa := []int{}
  for {
    if len(Ac) == 0 && len(Bc) == 0 {
      return msa
    } else if len(Ac) == 0 && len(Bc) != 0 {
      msa = append(msa, Bc[0])
      Bc = Bc[1:len(Bc)]
    } else if len(Ac) != 0 && len(Bc) == 0 {
      msa = append(msa, Ac[0])
      Ac = Ac[1:len(Ac)]
    } else {
      if Ac[0] < Bc[0] {
        msa = append(msa, Ac[0])
        Ac = Ac[1:len(Ac)]
      } else {
        msa = append(msa, Bc[0])
        Bc = Bc[1:len(Bc)]
      }
    }
  }
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
  
  _, A, _, B := LoadData("rosalind_mer.txt")
  msa := MergeSortedArrays(A, B)
  fmt.Println(IntSliceToString(msa))
}