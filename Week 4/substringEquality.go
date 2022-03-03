package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
  "math"
  "math/rand"
)

func precomputeHashes(s string, m, x int) []int {
  h := make([]int, len(s))
  h[0] = 0

  for i := 1; i < len(s); i++ {
    h[i] = ((h[i - 1] * x) + int(s[i])) % m
  }
  return h
}

func rangify(s string) []int {
  var temp []int
  for _, v := range strings.Split(s, " ") {
    convert, _ := strconv.Atoi(v)
    temp = append(temp, convert)
  }
  return temp
}

func solve(s string, ranges []string) {
  var (
    m1 = int(math.Pow10(9)) + 7
    m2 = int(math.Pow10(9)) + 9
    x1 = rand.Intn(int(math.Pow10(9))) + 1
    x2 = rand.Intn(int(math.Pow10(9))) + 1
    h1 = precomputeHashes(s, m1, x1)
    h2 = precomputeHashes(s, m2, x2)
  )

  for _, v := range ranges {
    c := rangify(v)
    i1 := c[0] + c[2]
    i2 := c[1] + c[2]

    if i1 >= len(s) {
      i1 = len(s) - 1
    }
    if i2 >= len(s) {
      i2 = len(s) - 1
    }

    if h1[i1] - (int(math.Pow(float64(x1), float64(c[2]))) * h1[c[0]]) % m1 == h1[i2] - (int(math.Pow(float64(x1), float64(c[2]))) * h1[c[1]]) % m1 && h2[i1] - (int(math.Pow(float64(x2), float64(c[2]))) * h2[c[0]]) % m2 == h2[i2] - (int(math.Pow(float64(x2), float64(c[2]))) * h2[c[1]]) % m2 {
      fmt.Println("Yes")
    } else {
      fmt.Println("No")
    }
  }
}

func main() {
  var (
    ranges    []string
    s         string
    numRanges int
  )
  fmt.Scanln(&s)
  fmt.Scanln(&numRanges)
  scanner := bufio.NewScanner(os.Stdin)

  for i := 0; i < numRanges; i++ {
    scanner.Scan()
    ranges = append(ranges, scanner.Text())
  }

  solve(s, ranges)
}
