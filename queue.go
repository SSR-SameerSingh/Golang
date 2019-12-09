package main

import "fmt"

type queue struct{
	elements []int
}

func (q1 *queue) enqueue(x int){
	q1.elements = append(q1.elements, x)
}

func (q1 *queue) dequeue()int{
	if len(q1.elements) == 0{
		fmt.Println("queue is empty, hence returning -1")
		return -1
	}else{
		val := q1.elements[0]
		q1.elements = q1.elements[1:]
		return val
	}
}

func (q1 *queue) getFront()int{
	if len(q1.elements) == 0{
		fmt.Println("Queue is empty, hence returning -1")
		return -1
	}else{
		return q1.elements[0]
	}
}


func (q1 *queue) isEmpty()bool{
	if len(q1.elements) == 0{
		return true
	}else{
		return false
	}
}


func main(){
	// q1 := queue{elements:[]int{}}
	q1 := new(queue)
	q1.enqueue(10)
	q1.enqueue(11)
	q1.enqueue(12)
	q1.enqueue(13)
	fmt.Println("Elements:",q1.elements)
	fmt.Println("Dequeue:",q1.dequeue())
	fmt.Println("Elements:",q1.elements)
	fmt.Println("isEmpty:",q1.isEmpty())
	fmt.Println("getFront:",q1.getFront())
}