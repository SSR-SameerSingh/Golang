package main

import "fmt"

type node struct{
	val int
	next *node
	prev *node
}

type dll struct{
	head *node
	size int
}

func (l *dll) init(){
	l.head = nil
	l.size = 0
}

func (l *dll) addNode(x int){
	l.size += 1
	new_node := &(node{val:x})
	if l.head == nil{
		l.head = new_node
	}else if l.head != nil{
		curr := l.head
		// prev := nil
		for curr.next != nil{
			curr = curr.next
		}
		curr.next = new_node
		new_node.prev = curr
	}
}

func (l *dll) removeNode(x int)bool{
	if l.head == nil{
		return false
	}else{
		curr := l.head
		for curr != nil{
			if curr.val != x{
				curr = curr.next
			}else if curr.val == x{
				l.size -= 1
				prev_node := curr.prev
				if prev_node == nil{
					l.head = curr.next
					return true
				}else if prev_node != nil{
					prev_node.next = curr.next
					return true
				}
			}
		}
		return false
	}
}

func (l *dll) length() int{
	return l.size
}

func (l *dll) print(){
	if l.head != nil{
		curr := l.head
		for curr != nil{
			fmt.Println(curr.val)
			curr = curr.next
		}
	}
}

// func (l *dll) 

func main(){
	dll := new(dll)
	dll.init()
	dll.addNode(10)
	dll.addNode(11)
	dll.addNode(12)
	dll.addNode(13)
	dll.removeNode(13)
	fmt.Println(dll.length())
	dll.print()
	// fmt.Println(dll.head.val)
	// fmt.Println("DLL")

}