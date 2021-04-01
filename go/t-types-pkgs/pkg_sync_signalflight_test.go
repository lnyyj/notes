package ttp

import (
	"fmt"
	"sync"
	"testing"

	"golang.org/x/sync/singleflight"
)

func Test_signalDO(t *testing.T) {
	sg := &singleflight.Group{}
	ch := make(chan struct{}, 0)
	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			<-ch
			ret, err, shard := sg.Do("11111", func() (v interface{}, err error) {
				// time.Sleep(500 * time.Millisecond)
				fmt.Println("->", i)
				return i, nil
			})
			wg.Done()
			fmt.Println("------>", ret, err, shard)
		}(i)
	}
	close(ch)
	wg.Wait()
}
