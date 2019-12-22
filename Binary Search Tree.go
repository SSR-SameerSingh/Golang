package main

import "fmt"

type node struct{
	val int
	left *node
	right *node
}

type bst struct{
	root *node
}

type stack struct{
	elements []node
}

func (s *stack) push(x node){
	s.elements = append(s.elements, x)
}

func (s *stack) pop()node{
	p := len(s.elements)-1
	temp := s.elements[p]
	s.elements = s.elements[:p]
	return temp
}

func (s *stack) isEmpty()bool{
	if len(s.elements) == 0{
		return true
	}else{
		return false
	}
}

// func (s *stack) top()node{
// 	if len(s.elements) > 0{
// 		return s.elements[-1]
// 	}
// }

func (b *bst) init(){
	b.root = nil
}

func (b *bst) add(x int){
	new_node := &(node{val:x})
	if b.root == nil{
		b.root = new_node
	}else{
		curr := b.root
		// s1 = new(stack)
		// s1.elements = append(s1.elements, curr)
		for curr != nil{
			if x <= curr.val{
				if curr.left != nil{
					curr = curr.left
				}else{
					curr.left = new_node
					break
				}
			}else if x > curr.val{
				if curr.right != nil{
					curr = curr.right
				}else{
					curr.right = new_node
					break
				}
			}
		}
	}
}



func (b *bst) inOrder()[]int{
	final := []int{}
	if b.root == nil{
		fmt.Println("Tree is empty!")
		return final
	}else if b.root != nil{
		// stack1 := new(stack)
		stack1 := stack{}
		curr := b.root
		// stack1.push(*curr)
		for stack1.isEmpty() != true || curr != nil{
			for curr != nil{
				// fmt.Println("push",curr.val)
				stack1.push(*curr)
				curr = curr.left
			}
			temp := stack1.pop()
			// fmt.Println("pop",temp.val)
			final = append(final, temp.val)
			// curr = temp.right
			if temp.right != nil{
				curr = temp.right
			}
		}
	}
	return final
}

func (b *bst) preOrder()[]int{
	final := []int{}
	if b.root == nil{
		fmt.Println("Tree is empty!")
		return final
	}else if b.root != nil{
		stack1 := new(stack)
		curr := b.root
		for stack1.isEmpty() != true || curr != nil{
			for curr != nil{
				final = append(final, curr.val)
				stack1.push(*curr)
				curr = curr.left
			}
			temp := stack1.pop()
			if temp.right != nil{
				curr = temp.right
				// final = append(final, curr.val)
			}
		}
	}
	return final
}

func main(){
	tree := new(bst)
	// s1 := new(stack)
	// s1.push(10)
	// order := tree.inOrder()
	// fmt.Println(order)
	tree.init()
	tree.add(9)
	tree.add(7)
	tree.add(19)
	tree.add(12)
	tree.add(1)
	tree.add(234)
	tree.add(10)
	// order := tree.inOrder()	
	inorder := tree.inOrder()
	fmt.Println("inOrder-->",inorder)
	preorder := tree.preOrder()
	fmt.Println("preOrder-->",preorder)
}