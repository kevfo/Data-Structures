package main

import (
  "fmt"
  "math/rand"
)

func areEqual(s1, s2 string) bool {
  if len(s1) != len(s2) {
    return false
  } else {
    for i := 0; i < len(s1); i++ {
      if s1[i] != s2[i] {
        return false
      }
    }
  }
  return true
}

func polyHash(s string, p, x int) int {
  hash := 0
  for i := len(s) - 1; i >= 0; i-- {
    hash = (hash * x + int(s[i])) % p
  }
  return hash
}

func precomputeHashes(text, pattern string, p, x int) []int {
  h := make([]int, len(text) - len(pattern) + 1)
  s := text[(len(text) - len(pattern)):len(text)]
  h[len(text) - len(pattern)] = polyHash(s, p, x)
  y := 1

  for i := 1; i <= len(pattern); i++ {
    y = (y * x) % p
  }
  for i := len(text) - len(pattern) - 1; i >= 0; i-- {
    h[i] = ((x * h[i + 1]) + int(text[i]) - (y * int(text[i + len(pattern)]))) % p
  }
  return h
}

func rabinKarp(text, pattern string) []int {
  p := 6700417  // "p" must be some large prime number!
  x := rand.Intn(p) + 1 
  var result []int
  pHash := polyHash(pattern, p, x)
  h := precomputeHashes(text, pattern, p, x)

  for i := 0; i <= len(text) - len(pattern); i++ {
    if pHash != h[i] {
      continue
    }
    if areEqual(text[i:(i + len(pattern))], pattern) {
      result = append(result, i)
    }
  }
  return result
}

func main() {
  var string, pattern string
  fmt.Scanln(&pattern)
  fmt.Scanln(&string)

  for _, v := range rabinKarp(string, pattern) {
    fmt.Printf("%v ", v)
  }
  fmt.Println()
}
