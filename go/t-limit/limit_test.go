package tlimit

import (
	"log"
	"testing"
	"time"

	"github.com/yangwenmai/ratelimit/simpleratelimit"
)

func Test_Limit(t *testing.T) {
	// rl := simpleratelimit.New(1, time.Second)
	// for i := 0; i < 100; i++ {
	// 	log.Printf("limit result: %v\n", rl.Limit())
	// }
	// log.Printf("limit result: %v\n", rl.Limit())

	rl := simpleratelimit.New(1, time.Second)
	for i := 0; i < 100; i++ {
		ok := rl.Limit()
		// time.Sleep(100 * time.Millisecond)
		log.Printf("limit result1: %d %v\n", i, ok)
		if ok {
			// rl.Undo()
			continue
		}
	}
	log.Printf("limit result2: %v\n", rl.Limit())
}
