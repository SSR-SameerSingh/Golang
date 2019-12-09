package main

import "fmt"

type node struct{
	val int
	next *node
}

type list struct{
	size int
	head *node
}

func (l *list) init(){
	l.head = nil
	l.size = 0
}

func (l *list) add(v int){
	l.size += 1
	new_node := &node{val:v}
	if l.head == nil{
		l.head = new_node
	}else{
		temp := l.head
		for temp.next != nil{
			temp = temp.next
		}
		temp.next = new_node
	}
}

func (l *list) print(){
	if l.head != nil{
		curr := l.head
		for curr != nil{
			fmt.Println(curr.val)
			curr = curr.next
		}
	}
}

func (l *list) remove_last(){
	if l.head != nil{
		curr := l.head
		prev := (*node)(nil)
		for curr.next != nil{
			prev = curr
			curr = curr.next
		}
		// curr.next = nil
		curr = nil
		prev.next = nil
		l.size -= 1
	}
}

func (l *list) remove_element(x int)bool{
	if l.head == nil{
		return false
	}else{
		curr := l.head
		prev := (*node)(nil)
		for curr != nil{
			if curr.val != x{
				prev = curr
				curr = curr.next
			}else if curr.val == x{
				if prev != nil{
					prev.next = curr.next
					l.size -= 1
					return true
				}else if prev == nil{
					l.head = curr.next
					l.size -= 1
					return true
				}
			}
		}
		return false
	}
}

func (l *list) length ()int{
	return l.size
}

func main(){
	ll := new(list)
	// ll.print()
	ll.init()
	ll.add(10)
	ll.add(11)
	ll.add(12)
	ll.add(13)
	ll.remove_last()
	ll.remove_element(12)
	ll.print()
}