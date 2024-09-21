package algos

import (
    "fmt"
    "adventofcode/utils"
)

const K = 3


var AhoTree = make([]TrieNode, 1)

type TrieNode struct {
    next[K] int
    link int
    p int //Parent node
    ch byte
    output int
}

func NewTrieNode() TrieNode {
    node := TrieNode {
        link: -1, 
        p: -1,
        output: -1,
        ch: 0,
    }
    for i := range node.next {
        node.next[i] = -1
    }
    return node
}

func NewAhoTree(str_set []string) []TrieNode {
    tree := []TrieNode{NewTrieNode()}
    for _, str := range str_set {
        tree = addString(tree, str)
    }
    tree = stateMachineify(tree)
    return tree
} //NewAhoTree

func addString(tree []TrieNode, str string) []TrieNode {
    if len(tree) < 1 {
        return tree
    }
    cur_node := 0
    for i := range str {
        ch_index := str[i] - 'a'
        if i == (len(str) - 1) { tree[cur_node].output = i }
        if tree[cur_node].next[ch_index] == -1 {
            new_node := NewTrieNode()
            new_node.p = cur_node
            new_node.ch = str[i]
            tree = append(tree, new_node)
            tree[cur_node].next[ch_index] = len(tree) - 1
            cur_node = len(tree) - 1
        } else {
            cur_node = tree[cur_node].next[ch_index]
        }
    }
    return tree
} //addString


func stateMachineify(tree []TrieNode) []TrieNode {
    if len(tree) < 1 { return tree; }
    visited := make([]bool, len(tree))
    queue := utils.Queue[int]{}
    queue.Enqueue(0)
    tree[0].link = 0 
    for !queue.IsEmpty() {
        curr, _ := queue.Dequeue()
        visited[curr] = true;
        cur_link := tree[curr].link 
        //fmt.Println(cur_link)
        for ch, next := range tree[curr].next { 
            if next != -1 {      
                if !visited[next] { 
                    queue.Enqueue(next)
                }
                if curr == 0 { 
                    tree[next].link = 0 
                } else {
                    tree[next].link = tree[cur_link].next[ch]
                } 
            } else {
                if curr == 0 {
                    tree[curr].next[ch] = 0
                } else {
                    tree[curr].next[ch] = tree[cur_link].next[ch]
                }
            } 
        } 
    }
    return tree
} //stateMachineify


func AhoPrettyPrint(tree []TrieNode) {
    if len(tree) < 1 { return; } 
    visited := make([]bool, len(tree))
    queue := utils.Queue[int]{}
    queue.Enqueue(0)  
    for !queue.IsEmpty() {
        curr, _ := queue.Dequeue() 
        visited[int(curr)] = true;  
        fmt.Print("(") 
        for i, next := range tree[curr].next {
            if (next != -1) && !visited[next] {
                fmt.Printf("%c, ", (i+'a'))
                queue.Enqueue(next)
            }
        }
        fmt.Print(")\n")
    }
}

func AhoUglyPrint(tree []TrieNode) {
    if len(tree) < 1 { return; }
    for i, node := range tree {
        for _, next := range node.next {
            fmt.Print(next, " ,") 
        }
        fmt.Printf(" link: %d \n", tree[i].link)
    }
}
