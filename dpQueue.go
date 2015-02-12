package dpQueue 

import ("errors")

type dpQueue struct{
	queue []dpQueueItem
}

type dpQueueItem struct{
	item interface{}
	priorityA float64
	priorityB float64
}

func New() *dpQueue{
	return new(dpQueue)
}

func (q *dpQueue) Enqueue(item interface{}, a float64, b float64){
	q.queue = append(q.queue, dpQueueItem{item, a, b})
}

func (q *dpQueue) DequeueA() interface{}{
	if(q.Count() == 0){ return errors.New("List is Empty") }
	a,_ := q.highPriorityIndices()
	return q.popItem(a)
}

func (q *dpQueue) DequeueB() interface{}{
	if(q.Count() == 0){ return errors.New("List is Empty") }
	_,b := q.highPriorityIndices()
	return q.popItem(b)
}

func (q *dpQueue) highPriorityIndices() (int, int){
	var highIndexA, highIndexB int
	var highPriorityA, highPriorityB float64
	for index, item := range q.queue{
		if item.priorityA > highPriorityA {
			highPriorityA, highIndexA = item.priorityA, index
		}
		if item.priorityB > highPriorityB {
			highPriorityB, highIndexB = item.priorityB, index
		}
	}
	return highIndexA, highIndexB
}

func (q *dpQueue) popItem(index int) interface{}{
	highItem := q.queue[index].item
	q.queue = append(q.queue[:index], q.queue[index + 1:]...)
	return highItem
}

func (q *dpQueue) Count() int{
	return len(q.queue)
}

func (q *dpQueue) Clear(){
	q.queue = make([]dpQueueItem,0)
}