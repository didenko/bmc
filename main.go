package main

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
    "strings"
)

type Node struct {
    L *Node
    R *Node
    Val string
}

var (
    root *Node
    depths map[*Node]uint32
)

const (
    LEFT = false
    RIGHT = true
)

func parse(s string) (id, l, r string) {

    parts := strings.Split(s, ",")
    id = parts[0]

    if len(parts) > 1 && parts[1] != "" {
        l = parts[1]
    }

    if len(parts) > 2 && parts[2] != "" {
        r = parts[2]
    }

    return id, l, r
}

func Root(at, l, r string) *Node {

    rt := &Node{nil, nil, at}
    depths[rt] = 0
    rt.Add(LEFT, l, 1)
    rt.Add(RIGHT, r, 1)
    return rt
}

func (n *Node) Add(side bool, id string, depth uint32) {

    if id != "" {
        newNode := &Node{nil, nil, id}
        depths[newNode] = depth

        if side == LEFT {
            n.L = newNode
        } else {
            n.R = newNode
        }
    }
}

func (n *Node) Insert(id, l, r string, depth uint32) {

    depth++

    if n.Val == id {
        n.Add(LEFT, l, depth)
        n.Add(RIGHT, r, depth)
        return
    }

    if n.L != nil{
        n.L.Insert(id, l, r, depth)
    }

    if n.R != nil{
        n.R.Insert(id, l, r, depth)
    }
}

func main() {

    depths = make(map[*Node]uint32)

    input := bufio.NewScanner(os.Stdin)

    for input.Scan() {
        if root == nil{
            root = Root(parse(input.Text()))
        } else {
            at, l, r := parse(input.Text())
            root.Insert(at, l, r, 0)
        }
    }

    if err:=input.Err(); err != nil{
        log.Fatalf("Errored: %v", err)
    }

    var shallow uint32 = math.MaxUint32

    for node, depth := range(depths){
        if node.L == nil && node.R == nil && depth < shallow {
            shallow = depth
        }
    }

    fmt.Println(shallow)
}
