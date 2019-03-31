package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (string, int, [][]float64) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  tb := string(b)
  content := []string{}
  scanner := bufio.NewScanner(strings.NewReader(tb))
  for scanner.Scan() {
    content = append(content, scanner.Text())
  }

  text := content[0]
  k, _ := strconv.Atoi(content[1])

  profile := [][]float64{}
  for _, s := range content[2:len(content)] {
    line := []float64{}
    num_strs := strings.Fields(s)
    for _, c := range num_strs {
      d, _ := strconv.ParseFloat(string(c), 64)
      line = append(line, d)
    }
    profile = append(profile, line)
  }

  return text, k, profile
}


func ProfileMostProbableKMmer(text string, k int, profile [][]float64) (string) {

  dict := map[string]int{
    "A": 0,
    "C": 1,
    "G": 2,
    "T": 3,
  }

  max_prob := 0.0
  pmp_kmer := ""
  for i := 0; i < len(text) - k; i++ {
    subtext := text[i:i + k]
    prob := 0.0
    for j, c := range subtext {
      sc := string(c)
      prob += profile[dict[sc]][j]
    }
    if prob > max_prob {
      max_prob = prob
      pmp_kmer = subtext
    }
  }

  return pmp_kmer
}


func main() {
  text, k, profile := LoadData("rosalind_ba2c.txt")
  pmp_kmer := ProfileMostProbableKMmer(text, k, profile)
  fmt.Println(pmp_kmer)
}