/*
  Solution:

  https://en.wikipedia.org/wiki/Longest_common_subsequence_problem

*/

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

  id := ""
  text := string(b)
  DNAs := []string{}
  content := make(map[string]string)
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    line_text := scanner.Text()
    if strings.Index(line_text, ">") > -1 {
      id = line_text[1:len(line_text)]
      content[id] = ""
    } else {
      if id != "" {
        content[id] += line_text
      }
    }
  }

  for _, v := range content {
    DNAs = append(DNAs, v)
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


func Backtrack(C [][]int, s, t string, i, j int) (string) {

  if i == 0 || j == 0 {
    return ""
  } else if s[i - 1] == t[j - 1] {
    return Backtrack(C, s, t, i - 1, j - 1) + string(s[i - 1])
  } else {
    if C[i][j - 1] > C[i - 1][j] {
      return Backtrack(C, s, t, i, j - 1)
    } else {
      return Backtrack(C, s, t, i - 1, j)
    }
  }
}


func main() {
  
  s, t := LoadData("rosalind_lcsq.txt")

  C := LCSLength(s, t)
  res := Backtrack(C, s, t, len(s), len(t))
  fmt.Println(res)
}