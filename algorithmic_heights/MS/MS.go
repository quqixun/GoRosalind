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


func MergeSort(array []int) ([]int) {

  if len(array) <= 1 {
    return array
  }

  mid := len(array) / 2

  p1 := MergeSort(array[0:mid])
  p2 := MergeSort(array[mid:len(array)])

  return MergeSortedArrays(p1, p2)
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


func main() {

  _, array := LoadData("rosalind_ms.txt")

  ms := MergeSort(array)
  res := IntSliceToString(ms)
  Write2File("ms_output.txt", []string{res})
}