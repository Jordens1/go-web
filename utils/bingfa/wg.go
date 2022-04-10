package bingfa

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func Word() {

	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start ")

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
				time.Sleep(time.Millisecond * 1)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
				time.Sleep(time.Millisecond * 1)

			}
		}

	}()

	fmt.Println("wait")

	wg.Wait()
	fmt.Println("Terminal")

}

var (
	counter int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func IncCounter() {

	runtime.GOMAXPROCS(1)

	wg.Add(2)

	go incCon(1)
	go incCon(2)

	wg.Wait()
	fmt.Printf("finall counter %d", counter)

}

func incCon(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()

		atomic.AddInt64(&counter, 1)
		value := counter

		value++
		counter = value
		// runtime.Gosched()

		mutex.Unlock()
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
	rand.Seed(time.Now().Unix())
}

func ChannelTest() {
	court := make(chan int)
	wg.Add(2)
	go player("Nadal", court)
	go player("Djokovic", court)
	court <- 1
	wg.Wait()

}

func player(name string, court chan int) {
	defer wg.Wait()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("player %s won \n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("palyer miss %s \n", name)
			close(court)
			return
		}

		fmt.Printf("player hit %s %d \n", name, ball)
		ball++
		court <- ball
	}
}

func NoCacheChannel() {
	baton := make(chan int)
	wg.Add(1)
	go Runner(baton)
	baton <- 1
	wg.Wait()

}

func Runner(baton chan int) {
	var newRunner int
	runner := <-baton
	fmt.Printf("Runner %d running with baton \n", runner)
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d to the line \n", runner)
		go Runner(baton)
	}
	time.Sleep(100 * time.Millisecond)
	if runner == 4 {
		fmt.Printf("Runner %d finish \n", runner)
		wg.Done()
		return
	}
	fmt.Printf("Runner %d exchange with runner %d \n", runner, newRunner)
	baton <- newRunner
}

const (
	numberGoroutines = 4
	taskLoad         = 10
)

func BufferChannel() {
	tasks := make(chan string, taskLoad)
	wg.Add(numberGoroutines)

	for gr := 0; gr < numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	for post := 0; post < taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	close(tasks)

	wg.Wait()

}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("worker : %d shutdown \n", worker)
			return
		}

		fmt.Printf("worker:%d , start : %s \n", worker, task)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("worker :%d , complete %s \n", worker, task)
	}
}
