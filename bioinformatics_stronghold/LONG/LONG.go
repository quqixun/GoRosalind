package main

import (
  "fmt"
  "math"
  "bufio"
  "strings"
  "io/ioutil"
)


func load_DNAs(file_path string) ([]string) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }
  text := string(b)

  id := ""
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

  DNAs := []string{}
  for _, v := range content {
    DNAs = append(DNAs, v)
  }

  return DNAs
}


func overlap(a, b string) ([]int) {

  // a: target, b: source
  len_a, len_b, min_len := len(a), len(b), 0
  if len_a < len_b {
    min_len = int(math.Ceil(float64(len_b) / 2.0))
  } else {
    min_len = int(math.Ceil(float64(len_a) / 2.0))
  }

  a_before_b, before := false, []int{-1, -1}
  a_after_b, after := false, []int{-1, -1}

  // if a before b
  for i := len_b; i >= min_len; i--{
    sub_b := b[0:i]
    if strings.HasSuffix(a, sub_b) {
      a_before_b = true
      before = []int{i - 1, len(sub_b)}
    }
  }

  // if a after b
  for i := 0; i <= len_b - min_len; i++ {
    sub_b := b[i:len_b]
    if strings.HasPrefix(a, sub_b) {
      a_after_b = true
      after = []int{0, len(sub_b)} 
    }
  }

  // 0 for before and 1 for after
  before_return := []int{0, before[0], before[1]}
  after_return := []int{1, after[0], after[1]}

  if a_before_b && a_after_b {
    if before[1] > after[1] {
      return before_return
    } else {
      return after_return
    }
  } else if a_before_b && !a_after_b {
    return before_return
  } else if !a_before_b && a_after_b {
    return after_return
  } else {
    return []int{-1, 0, 0}
  }
}


func is_start(dna_table [][]int) (bool) {

  has_0, has_1 := false, false
  for _, dt := range dna_table {
    if dt[0] == 0 {
      has_0 = true
    } else if dt[0] == 1 {
      has_1 = true
    } else {
      continue
    }
  }

  if has_0 && !has_1 {
    return true
  } else {
    return false
  }
}


func fragment_assemble(DNAs []string) (string) {

  ol_table := make([][][]int, 0)
  start := -1

  for i, dna1 := range DNAs {
    dna1_table := make([][]int, 0)
    for _, dna2 := range DNAs {
      if dna1 == dna2 {
        dna1_table = append(dna1_table, []int{-2, 0, 0})
        continue
      } else {
        dna1_table = append(dna1_table, overlap(dna1, dna2))
      }
    }

    if is_start(dna1_table) {
      start = i
    }

    ol_table = append(ol_table, dna1_table)
  }

  superstring := DNAs[start]
  next_idx := start
  for {
    if next_idx == -1 {
      break
    }

    next_dna_table := ol_table[next_idx]
    is_last := true
    for i, ndt := range next_dna_table {
      if ndt[0] == 0 {
        next_idx = i
        is_last = false
      }
    }

    next_str := ""
    next_DNA := DNAs[next_idx]
    next_str_idx := next_dna_table[next_idx][2]
    if !is_last {
      next_str = next_DNA[next_str_idx:len(next_DNA)]
    } else {
      next_str = next_DNA[len(next_DNA) - next_str_idx:len(next_DNA)]
      next_idx = -1
    }
    superstring += next_str
  }

  return superstring
}


func main() {
  
  DNAs := load_DNAs("rosalind_long.txt")
  fa := fragment_assemble(DNAs)
  fmt.Println(fa)
}