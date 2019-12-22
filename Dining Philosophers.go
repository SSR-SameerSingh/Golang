/* Dining Programmers (Dining Philosophers) problem

 
 1. Consider Every Programmer as independent machine, who decides when to eat and do what
 2. obtain locks on number of chopsticks (chop variable) whenever it is picked up, and unlock when it is done
 3. count variable to store the count of chopsticks each Programmer 
 4. Pickup method: if not enough chopstick are available, 
 	the goroutine( or thread) waits until they are available, using Conditional Variable 
 5. Done method: when a Programmer is done eating, it places chopsticks back, and 
 	signal construct is used to wake up the next waiting goroutine/thread in queue
 	FYI: using broadcast will wake every thread in queue, and might create deadlock or race condition

 */ 

package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func start(philo []string, index int, chop *int, count *[]int, cond *sync.Cond) {
	set_state(philo, index, "coding")
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	set_state(philo, index, "hungry")
	go pickup(philo, index, chop, count, cond)
	set_state(philo, index, "one")
	go pickup(philo, index, chop, count, cond)
	set_state(philo, index, "eating")
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	go done(philo, index, chop, count, cond)
	set_state(philo, index, "done")
	wg.Done()

}

func done(philo []string, index int, chop *int, count *[]int, cond *sync.Cond) {
	(*cond).L.Lock()
	fmt.Println(*count, *chop)
	if (*count)[index] > 0{
		(*count)[index] -= 2
		*chop += 2
		fmt.Println(*count, *chop)
		(*cond).Signal()
		(*cond).L.Unlock()
	}else{
		(*cond).L.Unlock()
	}
}

func pickup(philo []string, index int, chop *int, count *[]int, cond *sync.Cond) {
	(*cond).L.Lock()
	fmt.Println(*count, *chop)
	for *chop != 0 && (*count)[index] < 2{
		if (*chop) > 1 && (*count)[index] == 0 {
			(*count)[index] = 1
			*chop -= 1
			fmt.Println(*count, *chop)
		}else if (*chop) > 0 && (*count)[index] == 1{
			(*count)[index] = 2
			*chop -= 1
			fmt.Println(*count, *chop)
		}else{
			cond.Wait()
		}
	}
	(*cond).L.Unlock()
}
// coding, hungry, one, eating, done
func set_state(philo []string, index int, str string) {
	if str == "coding" {
		fmt.Printf("%s is %s. DO NOT DISTURB!\n", philo[index], str)
		time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	}
	if str == "hungry" {
		fmt.Printf("%s is feeling %s now.\n", philo[index], str)
		time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	}
	if str == "one" {
		fmt.Printf("%s found %s chopstick on table.\n", philo[index], str)
		time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	}
	if str == "eating" {
		fmt.Printf("%s found Second chopstick table %s. Now Eating, Do not DISTURB!\n", philo[index], str)
		time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	}
	if str == "done" {
		fmt.Printf("%s is %s eating. Putting both chopstick back at table!\n", philo[index], str)
		time.Sleep(time.Second * time.Duration(rand.Intn(10)))
		fmt.Printf("%s is going back to work.CODING. DO NOT DISTURB!\n", philo[index])
	}
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	philo := [] string{"P0","P1","P2","P3","P4"}
	count := [] int{0,0,0,0,0}
	chop := 5
	cond := sync.NewCond(&mutex)

	for i := 0; i < len(philo); i++ {
		wg.Add(1)
		go start(philo, i, &chop, &count, cond)
	}
	wg.Wait()

	fmt.Println("everybody is Finished eating.")
}