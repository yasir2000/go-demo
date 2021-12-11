package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type Thing struct {
	Uuids []string `json:uuids`
	mut   sync.Mutex
}

func (t *Thing) Gen(wg *sync.WaitGroup, c chan error) {
	defer wg.Done()

	uid, err := uuid.NewRandom()
	if err != nil {
		c <- err
		return
	}

	t.mut.Lock()
	t.Uuids = append(t.Uuids, fmt.Sprintf("%s", uid))
	t.mut.Unlock()

}

func Runner() {
	t := new(Thing)
	l := 1000

	c := make(chan error, l)

	wg := new(sync.WaitGroup)

	wg.Add(l)
	for i := 0; i < l; i++ {
		go t.Gen(wg, c)
	}

	if len(c) > 0 {
		for err := range c {
			fmt.Println(err)
		}
	}
}
