package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func load_sequence(file_path string) ([]int, bool) {

  sequence := []int{}

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
    return sequence, false
  }

  text := string(b)
  content := []string{}
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    content = append(content, scanner.Text())
  }

  sequence_chars := strings.Fields(content[1])
  for _, sc := range sequence_chars {
    sn, _ := strconv.Atoi(sc)
    sequence = append(sequence, sn)
  }

  return sequence, true
}


func init_slice(init_value int, length int) ([]int) {

  s := []int{}
  for i := 1; i<= length; i++ {
    s = append(s, init_value)
  }

  return s
}


func find_max_index(L []int) (int) {

  index := -1
  max_value := 0
  for i, l := range L {
    if l > max_value {
      max_value = l
      index = i
    }
  }

  return index
}


func reverse(sequence []int) ([]int) {

  if len(sequence) == 0 {
    return sequence
  }

  return append(reverse(sequence[1:]), sequence[0])
}


func longest_subsequence(sequence []int, mode string) ([]int) {

  L := init_slice(-1, len(sequence))
  P := init_slice(-1, len(sequence))

  if mode == "decrease" {
    sequence = reverse(sequence)
  }

  for i := 0; i < len(sequence); i++ {
    L[i] = 1
    for j := 0; j < i; j++ {
      if (sequence[j] < sequence[i]) && (L[j] + 1 > L[i]){
        P[i] = j
        L[i] = L[j] + 1
      }
    }
  }

  index := find_max_index(L)
  longeset := []int{sequence[index]}

  for {
    if P[index] == -1 {
      break
    } else {
      index = P[index]
      longeset = append(longeset, sequence[index])
    }
  }

  if mode == "increase"{
    longeset = reverse(longeset)  
  }
  return longeset
}


func print_sequence(sequence []int) () {

  sequence_chars := []string{}
  for _, n := range sequence {
    c := strconv.Itoa(n)
    sequence_chars = append(sequence_chars, c)
  }

  sequence_str := strings.Join(sequence_chars, " ")
  fmt.Println(sequence_str)
}


func main() {
  
  sequence, is_load_success := load_sequence("rosalind_lgis.txt")

  if is_load_success {
    li := longest_subsequence(sequence, "increase")
    ld := longest_subsequence(sequence, "decrease")
    
    print_sequence(li)
    print_sequence(ld)
  }
}