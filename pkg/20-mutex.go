/*Program with race condition

The output will be different for each time because of race condition. Some of the outputs which I encountered are final value of x 941,
final value of x 928, final value of x 922 and so on.
*/

package pkg

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}
func Mutex_race_condition_main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

//////////////////
/*Solving the race condition using mutex*/
var y = 0

func increment_for_mutex(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	y = y + 1
	m.Unlock()
	wg.Done()
}
func Mutex_race_condition_solving_by_mutex() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment_for_mutex(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of y", y)
}

///////////////////////////
/*Solving the race condition using channel*/
var z = 0

func increment_for_channel(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	z = z + 1
	<-ch
	wg.Done()
}
func Mutex_race_condition_solving_by_channel() {
	var w sync.WaitGroup
	ch := make(chan bool, 20)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment_for_channel(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of z", z)
}
