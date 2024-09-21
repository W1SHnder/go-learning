package utils

type Queue[T any] struct {
    items []T
}


func (q *Queue[T]) Enqueue(item T) {
    q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
    if q.IsEmpty() { 
        var zero T
        return zero, false
    }
    item := q.items[0]
    q.items = q.items[1:]
    return item, true
}

func (q *Queue[T]) IsEmpty() bool {
    return len(q.items) == 0
}

func (q *Queue[T]) Size() int {
    return len(q.items)
}

