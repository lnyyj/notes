package ttp

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Barrier struct {
	l    sync.Mutex
	cond *sync.Cond
}

func NewBarrier() *Barrier {
	b := &Barrier{}
	b.cond = sync.NewCond(&(b.l))
	// b.l.Lock()
	return b
}

func (b *Barrier) Broadcast() {
	fmt.Println("----->唤醒")
	b.cond.Broadcast()
}

func (b *Barrier) Wait(i int) {
	b.l.Lock()
	defer b.l.Unlock()
	b.cond.Wait()
	fmt.Println("----->", i)
}

func Test_synccond(t *testing.T) {
	b := NewBarrier()
	for i := 0; i <= 10; i++ {
		go b.Wait(i)
		if i == 10 {
			time.Sleep(1 * time.Second)
			b.Broadcast() // 唤醒全部等待
		}
	}
	select {}
}
