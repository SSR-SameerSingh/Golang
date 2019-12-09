package main

import "fmt"

type stack struct{
	elements []int
}

func (s1 *stack) push(x int){
	s1.elements = append(s1.elements, x)
}

func (s1 *stack) top()int{
	if len(s1.elements) == 0{
		fmt.Println("elements is empty, hence returning -1")
		return -1
	}
	p := len(s1.elements)-1
	return s1.elements[p]
}

func (s1 *stack) pop()int{
	if len(s1.elements) == 0{
		fmt.Println("No element to pop, hence returning -1")
		return -1
	}
	p := len(s1.elements)-1
	v := s1.elements[p]
	s1.elements = s1.elements[:p]
	return v
}

func (s1 *stack) isEmpty()bool{
	if len(s1.elements) == 0{
		return true
	}else{
		return false
	}
}

func (s1 *stack) size()int{
	return len(s1.elements)
}


func main(){
	// s1 := stack{elements : []int{}}
	s1 := new(stack)
	fmt.Println("Top element of stack:",s1.top())
	fmt.Println("All elements in stack:",s1.elements)
	s1.push(10)
	s1.push(11)
	s1.push(12)
	s1.push(13)
	fmt.Println("Top element of stack:",s1.top())
	(s1.pop())
	fmt.Println(s1.elements)
	fmt.Println("Top element of stack:",s1.top())
	fmt.Println("Size of stack:",s1.size())
	fmt.Println(s1.isEmpty())
}
