//simple topology sort with certain assumptions

package main

import (
    "fmt"
)

var print = fmt.Println

type Node struct {

    name string
    dependents []*Node
}

func CreateNode(nodeName string) *Node {
    
    return &Node{name: nodeName, dependents: []*Node{}}    
}

func (n *Node) AddDependent(node *Node) {

    n.dependents = append(n.dependents, node)
}

func (n *Node) AddDependents(nodes ...*Node) {

    for _, node := range nodes {
        n.dependents = append(n.dependents, node)
    }
}

type Graph struct {

    nodes []*Node   
}

func (g *Graph) AddNode(n *Node) {

   g.nodes = append(g.nodes, n)  
}

func TSort(n *Node, visitedNodes map[string]bool) {

    for _, d := range n.dependents {
        TSort(d, visitedNodes)
    }

    if _, ok := visitedNodes[n.name]; ok {
       return 
    }
    visitedNodes[n.name] = true
    print(n.name)
}

func (g *Graph) TopologySort(startNode *Node) {

    visitedNodes := make(map[string]bool, len(g.nodes))
    TSort(startNode, visitedNodes)
}

func main() {

    g := &Graph{nodes: []*Node{}}

    a := CreateNode("A")
    b := CreateNode("B")
    c := CreateNode("C")
    d := CreateNode("D")
    e := CreateNode("E")
    f := CreateNode("F")

    g.AddNode(a)
    g.AddNode(b)
    g.AddNode(c)
    g.AddNode(d)
    g.AddNode(e)
    g.AddNode(f)

    a.AddDependents(b, e)
    b.AddDependents(d, c)
    c.AddDependents(f)
    f.AddDependents(e)

    g.TopologySort(a)
}
