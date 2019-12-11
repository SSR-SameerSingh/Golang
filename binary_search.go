package main

import "fmt"

func binary_search(arr []int, start int, end int, target int)int{
	if start > end || start > len(arr){
		return -1
	}else if start == end{
		if arr[start] == target{
			return start
		}else{
			return -1
		}
	}
	mid := (start+end)/2
	if arr[mid] == target{
		return mid
	}else if arr[mid] < target{
		return binary_search(arr, mid+1, end, target)
	}else{
		return binary_search(arr, start, mid-1, target)
	}
}

func main(){
	arr := []int{0,1,2,3,4,5,6,7,8,9,10,11,12}
	fmt.Println(binary_search(arr, 0, len(arr)-1, 100))	
}