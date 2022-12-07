package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type node struct {
  name string
  children []*node
  size int
}

type fileTree struct {
  root *node
}

func parseInput() fileTree {
  var tree fileTree
  stack := []*node{}
  sc := bufio.NewScanner(os.Stdin)

  for sc.Scan() {
    line := sc.Text()
    
    // command
    if line[0] == '$' {
      if strings.Contains(line, "cd") {
        lineSplit := strings.Split(line, " ")
        dir := lineSplit[len(lineSplit)-1]
        
        if dir == ".." {
          stack = stack[:len(stack) - 1]
          continue
        }
        
        if dir == "/" {
          tree = fileTree{}
          tree.root = &node{dir, []*node{}, 0}
          stack = append(stack, tree.root) 
        } else {
          node := &node{dir, []*node{}, 0}
          head := stack[len(stack) - 1]
          head.children = append(head.children, node)
          stack = append(stack, node) 
        }
        
      } 
    } else {
      if line[0] == 'd' {
        continue
      }

      lineSplit := strings.Split(line, " ")
      size, _ := strconv.Atoi(lineSplit[0])
      nameFile := lineSplit[1]

      node := &node{nameFile, nil, size}
      head := stack[len(stack) - 1]
      head.children = append(head.children, node)
    }
  }

  return tree
}

func calcTotalSize(t *node, res *int) {
  if t == nil {
    return
  }

  for _, ch := range t.children {
    calcTotalSize(ch, res)
    t.size += ch.size
  }

  if t.children != nil && t.size < 100000{
    *res += t.size
  }
}

func partOne(t fileTree) {
  res := 0
  calcTotalSize(t.root, &res)
  fmt.Println(res)
}

func findDirSizes(t *node, sizes *[]int) {
  if t == nil {
    return
  }

  for _, ch := range t.children {
    if ch.children != nil {
      *sizes = append(*sizes, ch.size)
      findDirSizes(ch, sizes)
    }
  }
}

func partTwo(t fileTree) {
  res := 0
  calcTotalSize(t.root, &res)
  fmt.Println("total size", t.root.size)

  sizeArr := []int{}
  findDirSizes(t.root, &sizeArr)
  sort.Ints(sizeArr)

  // fmt.Println(sizeArr)

  remaining := 70000000 - t.root.size
  for _, s := range sizeArr {
    if s + remaining >= 30000000 {
      fmt.Println(s)
      break
    }
  }
}

func main() {
  fTree := parseInput();
  // fmt.Println(fTree.root)
  // partOne(fTree)
  partTwo(fTree)
}
