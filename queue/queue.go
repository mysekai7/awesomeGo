package queue

type Queue []interface{} //将int改为interface{}，支持任何类型，实现可用.()做类型限定

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int) //类型强转
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
