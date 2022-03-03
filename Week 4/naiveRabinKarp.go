package main

import (
  "fmt"
  "math/rand"
)

func polyHash(s string, p, x int) int {
  hash := 0
  for i := len(s) - 1; i >= 0; i-- {
    hash = ((x * hash) + int(s[i])) % p
  }
  return hash
}

func naiveRabinKarp(text, pattern string) []int {
  p := 6700417
  x := rand.Intn(p - 1) + 1
  var result []int
  pHash := polyHash(pattern, p, x)

  for i := 0; i <= len(text) - len(pattern); i++ {
    tHash := polyHash(text[i:(i + len(pattern))], p, x)
    if pHash != tHash {
      continue
    }
    if text[i:(len(pattern) + i)] == pattern {
      result = append(result, i)
    }
  }
  return result
}

func main() {
  var text, pattern string
  fmt.Scanln(&pattern)
  fmt.Scanln(&text)

  for _, v := range naiveRabinKarp(text, pattern) {
    fmt.Printf("%v ", v)
  }
  fmt.Println()
}
