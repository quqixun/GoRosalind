package main

import (
  "fmt"
  "math"
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
  samples := make(map[string]string)

  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    line_text := scanner.Text()
    if strings.Index(line_text, ">") > -1 {
      id = line_text[1:]
      samples[id] = ""
    } else {
      if id != "" {
        samples[id] += line_text
      }
    }
  }

  keys := []string{}
  for k := range samples {
    keys = append(keys, k)
  }

  s1 := samples[keys[0]]
  s2 := samples[keys[1]]
  return s1, s2
}


func InitDP(m, n int) ([][]int) {

  dp := [][]int{}
  for i := 0; i <= m; i++ {
    row := []int{}
    for j := 0; j <= n; j++ {
      row = append(row, 0)
    }
    dp = append(dp, row)
  }

  return dp
}


func Min3(x, y, z int) (int) {

  xf := float64(x)
  yf := float64(y)
  zf := float64(z)

  return int(math.Min(xf, math.Min(yf, zf)))
}


func EditDistance(s1, s2 string, m, n int) (int) {

  dp := InitDP(m, n)

  for i := 0; i <= m; i++ {
    for j := 0; j <= n; j++ {

      if i == 0 {
        dp[i][j] = j
      } else if j == 0 {
        dp[i][j] = i
      } else if s1[i - 1] == s2[j - 1] {
        dp[i][j] = dp[i - 1][j - 1]
      } else {
        dp[i][j] = 1 + Min3(dp[i][j - 1],
                            dp[i - 1][j],
                            dp[i - 1][j - 1])
      }
    }
  }

  return dp[m][n]
}


func main() {

  s1, s2 := LoadData("rosalind_edit.txt")
  ed := EditDistance(s1, s2, len(s1), len(s2))
  fmt.Println(ed)
}