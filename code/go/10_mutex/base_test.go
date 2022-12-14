package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

type user struct{}
type userMap map[int]*user

func TestUserMapLock(t *testing.T) {
	runtime.GOMAXPROCS(2)

	mutex := sync.RWMutex{}
	users := userMap{}

	c1GetLockCount := 0
	c2GetLockCount := 0
	c2Mutext := sync.Mutex{}

	// fill base user
	initUser := 500
	for i := 0; i < initUser; i++ {
		users[i] = &user{}
	}

	go func() {
		for i := 0; i < initUser; i++ {
			t1 := time.Now()

			mutex.Lock()
			delete(users, i)
			mutex.Unlock()

			cost := time.Now().Sub(t1).Microseconds()
			fmt.Printf("g1 get lock cost:%d us \n", cost)

			c1GetLockCount++
			time.Sleep(time.Duration(rand.Int()%100+100) * time.Millisecond)
		}
	}()

	consume := func(msgCount int, workNeedTime int, fetchTime time.Duration) {
		for {
			// 实际的批量消息中间的处理时间很可能没有间隔
			for i := 0; i < msgCount; i++ {
				// 理论上，如果是公平的，则g2持有一次锁 500用户数 * workNeedTime = 5ms，g1应该每5ms拿到一次锁
				mutex.RLock()
				for range users {
					// 每个耗时不一样 us
					time.Sleep(time.Duration(rand.Int()%workNeedTime+1) * time.Microsecond)
				}
				mutex.RUnlock()

				c2Mutext.Lock()
				c2GetLockCount++
				c2Mutext.Unlock()
			}
			time.Sleep(fetchTime)
		}
	}

	for p := 0; p < 1; p++ {
		go func() {
			consume(200, 20, time.Millisecond*100)
		}()
	}

	time.Sleep(time.Second * 10)
	fmt.Printf("g1 lock: %d, g2 lock:%d \n", c1GetLockCount, c2GetLockCount)
}

func Test2(t *testing.T) {
	runtime.GOMAXPROCS(3)
	var wg sync.WaitGroup
	const runtime = 1 * time.Second
	var sharedLock sync.Mutex
	greedyWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}
	politeWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("Polite worker was able to execute %v work loops\n", count)
	}
	wg.Add(2)
	go greedyWorker()
	go politeWorker()
	wg.Wait()
}

func BenchmarkLock(b *testing.B) {
	m1 := sync.Mutex{}
	send := func() {
		m1.Lock()
		time.Sleep(time.Microsecond * 1)
		m1.Unlock()
	}

	for i := 0; i < b.N; i++ {
		send()
	}
}

func BenchmarkLockAndRecover(b *testing.B) {
	m2 := sync.Mutex{}
	callFunc := func() {
		defer func() { recover() }()
		time.Sleep(time.Microsecond * 1)
	}

	send2 := func() {
		m2.Lock()
		defer m2.Unlock()
		callFunc()
	}

	for i := 0; i < b.N; i++ {
		send2()
	}
}

func TestRecover(t *testing.T) {
	c := make(chan interface{}, 0)
	close(c)

	send := func(num int) {
		defer func() {
			recover()
			fmt.Printf("recover")
		}()

		if num > 2 {
			fmt.Printf("write to closed channel\n")
			select {
			case c <- 66:
				fmt.Printf("write to closed channel 2222\n")
			default:
				fmt.Printf("default")
			}
		}
	}

	for i := 1; i <= 5; i++ {
		send(i)
		fmt.Println(i)
	}
}
