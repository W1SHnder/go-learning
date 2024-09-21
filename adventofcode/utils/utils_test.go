package utils

import "testing"

func TestQueue(t *testing.T) {
    strs := []string{"one", "two", "three", "four", "five"}
    queue := Queue{}
    var i = 0
    for _, str := range strs {
        if queue.Size() != i {
            t.Errorf("Failed enqueue on index %d\n", i)
        }
        queue.Enqueue(str)
        i += 1
    }
    i = 0 
    for {
        str := queue.Dequeue()
        if str == nil {
            if i < (len(strs) - 1) {
                t.Errorf("Failed deque (empty early) on index %d\n", i) 
            }
            return
        } 
        if str != strs[i] {
            t.Errorf("Failed dequeue on index %d\n", i)
        }
        i += 1
    }
}
