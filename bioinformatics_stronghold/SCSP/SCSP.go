package main

import (
  "fmt"
  "bufio"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (string, string) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  text := string(b)
  DNAs := []string{}
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    DNAs = append(DNAs, scanner.Text())
  }

  return DNAs[0], DNAs[1]
}


func Max(a, b int) (int) {
  if a > b {
    return a
  } else {
    return b
  }
}


func LCSLength(s, t string) ([][]int) {

  m, n := len(s), len(t)
  C := [][]int{}
  for i := 0; i <= m; i++ {
    Ci := []int{}
    for j := 0; j <= n; j++ {
      Ci = append(Ci, 0)
    }
    C = append(C, Ci)
  }

  for i := 1; i <= m; i++ {
    for j := 1; j <= n; j++ {
      if s[i - 1] == t[j - 1] {
        C[i][j] = C[i - 1][j - 1] + 1
      } else {
        C[i][j] = Max(C[i][j - 1], C[i - 1][j])
      }
    }
  }

  return C
}


func Backtrack(C [][]int, s, t string, i, j int, sidx, tidx *[]int) (string) {

  if i == 0 || j == 0 {
    return ""
  } else if s[i - 1] == t[j - 1] {
    *sidx = append([]int{i - 1}, *sidx...)
    *tidx = append([]int{j - 1}, *tidx...)
    return Backtrack(C, s, t, i - 1, j - 1, sidx, tidx) + string(s[i - 1])
  } else {
    if C[i][j - 1] > C[i - 1][j] {
      return Backtrack(C, s, t, i, j - 1, sidx, tidx)
    } else {
      return Backtrack(C, s, t, i - 1, j, sidx, tidx)
    }
  }
}


func CommonSupersequence(s, t, m string, sidx, tidx []int) (string) {

  cs := s[0:sidx[0]] + t[0:tidx[0]]

  for i := 0; i < len(m); i++ {
    cs += string(m[i])

    if i != len(m) - 1 {
      cs += s[sidx[i] + 1:sidx[i + 1]] + t[tidx[i] + 1:tidx[i + 1]]
    }
  }

  cs += s[sidx[len(sidx) - 1] + 1:len(s)] + t[tidx[len(tidx) - 1] + 1:len(t)]
  return cs
} 


func main() {
  
  s, t := LoadData("rosalind_scsp.txt")

  C := LCSLength(s, t)
  sidx, tidx := []int{}, []int{}
  m := Backtrack(C, s, t, len(s), len(t), &sidx, &tidx)
  cs := CommonSupersequence(s, t, m, sidx, tidx)
  fmt.Println(cs)
}