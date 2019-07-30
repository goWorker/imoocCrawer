package queue
// An FIFO queue
type Queue []int

//Pushs the element into the queue
func (q *Queue) Push(v int){
	*q = append(*q,v)
}


//Pops element from head
func (q *Queue) Pop() int{
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// Returns if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}


