package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
  "strings"
)

type Vertex struct {
  key     int
  sum     int
  left    *Vertex
  right   *Vertex
  parent  *Vertex
}

func update(v *Vertex) {
  if v == nil {return}

  if v.left == nil && v.right == nil {
    v.sum = v.key
  } else if v.left == nil {
    v.sum = v.key + v.right.sum
  } else if v.right == nil {
    v.sum = v.key + v.left.sum
  } else {
    v.sum = v.key + v.left.sum + v.right.sum
  }

  if v.left != nil {v.left.parent = v}
  if v.right != nil {v.right.parent = v}
}

func smallRotation(v *Vertex) {
  var parent, grandparent, m *Vertex = v.parent, v.parent.parent, nil
  if parent == nil {return}

  if parent.left == v {
    m = v.right
    v.right = parent
    parent.left = m
  } else {
    m = v.left
    v.left = parent
    parent.right = m
  }

  update(parent) ; update(v)
  v.parent = grandparent

  if grandparent != nil {
    if grandparent.left == parent {
      grandparent.left = v
    } else {
      grandparent.right = v
    }
  }
}

func bigRotation(v *Vertex) {
  if v.parent.parent.left == v.left && v.parent.left == v {
    smallRotation(v.parent) ; smallRotation(v)
  } else if v.parent.parent.right == v.right && v.parent.right == v {
    smallRotation(v.parent) ; smallRotation(v)
  } else {
    smallRotation(v) ; smallRotation(v)
  }
}

func splay(v *Vertex) *Vertex {
  if v == nil {return nil}

  for v.parent != nil {
    if v.parent.parent == nil {
      smallRotation(v)
      break
    }
    bigRotation(v)
  }
  return v
}

func find(root *Vertex, key int) (*Vertex, *Vertex) {
  var v, last, next *Vertex = root, root, nil

  for v != nil {
    if v.key >= key && (next == nil || v.key < next.key) {
      next = v
    }
    last = v
    if v.key == key {break}
    if v.key < key {
      v = v.right
    } else {
      v = v.left
    }
  }

  root = splay(last)
  if next != nil {root = splay(next)}
  return next, root
}

func split(root *Vertex, key int) (*Vertex, *Vertex) {
  result, root := find(root, key)
  if result == nil {return root, nil}
  right := splay(result)
  left := right.left
  right.left = nil

  if left != nil {left.parent = nil}

  update(left) ; update(right)
  return left, right
}

func merge(left, right *Vertex) *Vertex {
  if left == nil {return right}
  if right == nil {return left}

  for right.left != nil {
    right = right.left
  }

  right = splay(right) ; right.left = left ; left.parent = right
  update(right)
  return right
}

var root *Vertex = nil

func insert(x int) {
  left, right := split(root, x)
  var newVertex *Vertex = nil
  if right == nil || right.key != x {
    newVertex = &Vertex{key: x, sum: x, left: nil, right: nil, parent: nil}
  }
  root = merge(merge(left, newVertex), right)
}

func erase(x int) {
  left, middle := split(root, x)
  middle, right := split(middle, x + 1)
  root = merge(left, right)
}

func search(x int) bool {
  result, _ := find(root, x)
  if result == nil || result.key != x {
    return false
  } else {
    return true
  }
}

func sum(from, to int) int {
  left, middle := split(root, from)
  middle, right := split(middle, to + 1)
  ans := 0

  if middle != nil {
    ans += middle.sum
  }
  root = merge(merge(left, middle), right)

  return ans
}

const MODULO = 1000000001

func main() {
  var (
    numCommands     int
    commands        []string
    last_sum        int
  )
  fmt.Scanln(&numCommands)

  scanner := bufio.NewScanner(os.Stdin)

  for i := 0; i < numCommands; i++ {
    scanner.Scan()
    commands = append(commands, scanner.Text())
  }

  for _, v := range commands {
    command := strings.Split(v, " ")
    n1, _ := strconv.Atoi(command[1])
    switch command[0] {
    case "+":
      insert((n1 + last_sum) % MODULO)
    case "-":
      erase((n1 + last_sum) % MODULO)
    case "?":
      if search((n1 + last_sum) % MODULO) {
        fmt.Println("Found")
      } else {
        fmt.Println("Not found")
      }
    case "s":
      n2, _ := strconv.Atoi(command[2])
      res := sum((n1 + last_sum) % MODULO, (n2 + last_sum) % MODULO)
      fmt.Println(res)
      last_sum = res % MODULO
    }
  }
}
