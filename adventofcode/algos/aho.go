package algos

import (
    "fmt"
    "adventofcode/utils"
)

const K = 255


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
    for i, str := range str_set {
        tree = addString(tree, str, i)
    }
    tree = stateMachineify(tree)
    return tree
} //NewAhoTree

func addString(tree []TrieNode, str string, out_val int) []TrieNode {
    if len(tree) < 1 {
        return tree
    }
    cur_node := 0
    for i := range str {
        ch_index := str[i] //- 'a'
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
    tree[cur_node].output = out_val
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


func MatchStrings(tree []TrieNode, str string) []int {
    v := 0 
    var out []int
    for i := range str {
        //fmt.Printf("Checking: %c at %d; found? ", str[i], v)
		link := tree[v].link
		if (tree[v].output > -1) {
            //fmt.Print("yup")
            out = append(out, tree[v].output) 
        }
		if (tree[link].output > -1) {
			//fmt.Print("yup")
			out = append(out, tree[link].output)
		}
        //fmt.Print("\n")
        v = tree[v].next[str[i]] // - 'a'] 
    }
	if (tree[v].output > -1) {
		out = append(out, tree[v].output)
	}
    return out
} //MatchStrings

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
		fmt.Print(i, ": ")
        for _, next := range node.next {
            fmt.Print(next, ", ") 
        }
        fmt.Printf(" pchar: %c, link: %d, out: %d \n", tree[i].ch, tree[i].link, tree[i].output)
    }
}
