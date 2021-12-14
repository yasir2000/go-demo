package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	const gr = 100

	var wg sync.WaitGroup

	wg.Add(gr * 2)

	var n int = 0

	//1.
	var m sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			//2.
			m.Lock()
			n++
			//3.
			m.Unlock()
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			//2.
			m.Lock()
			//3.
			defer m.Unlock()
			n--
			//3.
			//m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Final value of n:", n)

}
