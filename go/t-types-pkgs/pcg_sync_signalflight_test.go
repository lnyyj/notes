package ttp

import (
	"fmt"
	"github.com/lnyyj/dir_links/gosrc/sync"
	"time"
	"golang.org/x/sync/singleflight"
	"testing"
)



func Test_signalDO(t *testing.T) {
 	sg := &singleflight.Group{}
 	ch := make(chan struct{},0)
 	wg := &sync.WaitGroup{}
 	wg.Add(100)
 	for i:= 0; i<100; i++ {
 		go func(i int) {
 			<-ch
 			sg.Do("11111", func() {
 				time.Sleep(500 * time.Millisecond)
 				fmt.Println("---->", i)
			})
 			wg.Done()
		}(i)
	}
 	close(ch)
 	wg.Wait()
}
