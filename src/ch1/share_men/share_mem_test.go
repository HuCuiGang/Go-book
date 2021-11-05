package share_men

import (
	"sync"
	"testing"
	"time"
)

func TestCounteer(t *testing.T)  {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d",counter)
}

func TestCounterThreadSafe(t *testing.T)  {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}

	time.Sleep(1 * time.Second)
	t.Logf("counter = %d",counter)
}

//互斥锁
func TestCounterWaitGrop(t *testing.T)  {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
				wg.Done()
			}()
			mut.Lock()
			counter++
			if counter == 10 {
				return
			}

		}()
	}
	wg.Wait()
	//time.Sleep(1 * time.Second)
	t.Logf("counter = %d",counter)
}

//读写锁
func TestCounterRWaitGrop(t *testing.T)  {
	var mut sync.RWMutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.RUnlock()
				wg.Done()
			}()
			mut.RLock()
			counter++
			if counter == 10 {
				return
			}

		}()
	}
	wg.Wait()
	//time.Sleep(1 * time.Second)
	t.Logf("counter = %d",counter)
}