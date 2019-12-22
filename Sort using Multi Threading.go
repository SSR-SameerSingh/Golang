/* Question:
	The goal of this activity is to explore the use of threads by creating a program 
	for sorting integers that uses four goroutines to create four sub-arrays 
	and then merge the arrays into a single array.

Approach:
	Comparison based sorting using threads
*/

package main

import (
	"fmt"
	"sort"
	"sync"
	// "runtime"
	// "copy"
	// "math"
	// "time"
)

var wg, wg1 sync.WaitGroup
// var wg1 sync.WaitGroup
var mutex sync.Mutex // doing controlled concurrency/synchronization, to prevent interleaving of shared variable


func send_main(arr [][]float64, cc chan float64) { // driver function
	// defer close(cc)
	for i:= range arr{
		wg1.Add(1)
		go send(arr[i], cc)
	}
	wg1.Wait()
	close(cc)
}

func send(arr []float64, cc chan<- float64) { // sending each element over channel
	for i:= range arr{
		cc <- arr[i]
	}
	wg1.Done()
}

func sorted(arr *[]float64) {
	defer wg.Done()
	sort.Float64s(*arr)
}

func insert(arr *[]float64, value float64) { // concurrently inserting the element in slice
	mutex.Lock() // to avoid race condition
	if len(*arr) == 0{
		// mutex.Lock() // no other go routine can access this section while mutex is locked. One go routine at a time can access arr shared variable.
		*arr = append(*arr, value)
		wg.Done()
		mutex.Unlock()
		return
	}else{
		// mutex.Lock() //this  will give race condition
		for i:=0;i<len(*arr);i++ {
			if (*arr)[i] < value { // putting incoming element in right spot in array
				continue
			}else if (*arr)[i] >= value {
				rem := (*arr)[i:]
				(*arr) = (*arr)[:i]
				temp := [] float64{}
				for j:= range rem{
					temp = append(temp, rem[j])
				}
				(*arr) = append(*arr, value)
				for j := range temp{
					(*arr) = append(*arr, temp[j])
				}
				wg.Done()
				mutex.Unlock()
				return
			}
		}
		(*arr) = append(*arr, value)
		wg.Done()
		mutex.Unlock()
		return
	}
}

func main() {

	// runtime.GOMAXPROCS(100)

	cc := make(chan float64, 100)
	final := []float64{}

	arr1 := []float64 {2,3,4,6,56,321,35,7,68,0,1,9}
	arr2 := []float64 {23,343,134,43,12,34,57,472,68,9}
	arr3 := []float64 {212,33,334,6,378,31,35,7,638,339}
	arr4 := []float64 {2,36,40,968,456,21,35,7,68,9,0}

	arr := [][]float64{arr1, arr2, arr3, arr4}

	// wg.Add(1)
	go send_main(arr, cc)
	// wg.Wait()
	for j := range cc{ // receiving each element via channel
		// final = append(final, j)
		wg.Add(1)
		// sorted(&final)
		go insert(&final, j) //  running go routine to insert the element at correct position in final slice
		// fmt.Println("arr",final)
	}
	wg.Wait()
	fmt.Println(final)
	fmt.Println(len(final))
}