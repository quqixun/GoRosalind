package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (string, int, int, int) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  text := string(b)
  content := []string{}
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    content = append(content, scanner.Text())
  }

  genome := content[0]
  num_str := content[1]
  num_strs := strings.Fields(num_str)

  k, _ := strconv.Atoi(num_strs[0])
  L, _ := strconv.Atoi(num_strs[1])
  t, _ := strconv.Atoi(num_strs[2])

  return genome, k, L, t
}


func IsElementExist(str_slice []string, str string) (bool) {

  for _, s := range str_slice {
    if str == s {
      return true
    }
  }

  return false
}


func FindLtClump(genome string, k, L, t int) ([]string) {

  Lt_clump := []string{}

  len_genome := len(genome)
  for i := 0; i <= len_genome - L; i++ {
    sub_genome := genome[i:i + L]
    sub_k_mers := make(map[string]int)
    for j := 0; j <= len(sub_genome) - k; j++ {
      sub_sub_genomre := sub_genome[j:j + k]
      sub_k_mers[sub_sub_genomre] += 1
    }

    for key, value:= range sub_k_mers {
      if value == t {
        if !IsElementExist(Lt_clump, key) {
          Lt_clump = append(Lt_clump, key)        
        }
      }
    }
  }

  return Lt_clump
}


func main() {
  
  genome, k, L, t := LoadData("rosalind_ba1e.txt")
  Lt_clump := FindLtClump(genome, k, L, t)
  fmt.Println(strings.Join(Lt_clump, " "))
}