package main

import (
  "fmt"
  "math"
  "bufio"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (string) {

  RNA := ""

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  scanner := bufio.NewScanner(strings.NewReader(string(b)))
  for scanner.Scan() {
    if strings.Index(scanner.Text(), ">") == -1 {
      RNA += scanner.Text()
    }
  }

  return RNA
}


func MotzkinNumberSum(RNA string, seqs map[string]int) (int) {

  if len(RNA) <= 1 {
    return 1
  }

  if _, ok := seqs[RNA]; ok {
    return seqs[RNA]
  }

  seqs[RNA] = MotzkinNumberSum(RNA[1:len(RNA)], seqs)
  for i := 1; i < len(RNA); i++ {
    RNA0 := string(RNA[0])
    RNAi := string(RNA[i])
    if (RNA0 == "A" && RNAi == "U") ||
       (RNA0 == "U" && RNAi == "A") ||
       (RNA0 == "C" && RNAi == "G") ||
       (RNA0 == "G" && RNAi == "C") {
      seqs[RNA] += MotzkinNumberSum(RNA[1:i], seqs) * 
                   MotzkinNumberSum(RNA[i + 1:len(RNA)], seqs)
    }
  }


  seqs[RNA] = int(math.Mod(float64(seqs[RNA]), 1000000.0))
  return seqs[RNA]
}


func main() {

  RNA := LoadData("rosalind_motz.txt")

  seqs := make(map[string]int)
  mns := MotzkinNumberSum(RNA, seqs)
  fmt.Println(mns)
}