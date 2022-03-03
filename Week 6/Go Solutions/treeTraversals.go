// Takes too long (fails on test case 16)

package main

import (
  "fmt"
)

type Node struct {
  numNodes    int
  lefts       []int
  rights      []int
  keys        []int
}

func (n *Node) read() {
  var a, b, c int

  fmt.Scanln(&n.numNodes)

  n.lefts = make([]int, n.numNodes)
  n.rights = make([]int, n.numNodes)
  n.keys = make([]int, n.numNodes)

  for i := 0; i < n.numNodes; i++ {
    fmt.Scan(&a) ; fmt.Scan(&b) ; fmt.Scan(&c)
    n.keys[i] = a ; n.lefts[i] = b; n.rights[i] = c
  }
}

func (n *Node) inOrderTraversal(root int) {
  if root == -1 {return}
  n.inOrderTraversal(n.lefts[root])
  fmt.Print(n.keys[root], " ")
  n.inOrderTraversal(n.rights[root])
}

func (n *Node) preOrderTraversal(root int) {
  if root == -1 {return}
  fmt.Print(n.keys[root], " ")
  n.preOrderTraversal(n.lefts[root])
  n.preOrderTraversal(n.rights[root])
}

func (n *Node) postOrderTraversal(root int) {
  if root == -1 {return}
  n.postOrderTraversal(n.lefts[root])
  n.postOrderTraversal(n.rights[root])
  fmt.Print(n.keys[root], " ")
}

func main() {
  var tree Node
  tree.read()
  tree.inOrderTraversal(0)
  fmt.Println("")
  tree.preOrderTraversal(0)
  fmt.Println("")
  tree.postOrderTraversal(0)
  fmt.Println("")
}
