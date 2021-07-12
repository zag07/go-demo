package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	stop1()
	stop2()
	stop3()
	stop4()
}

func stop1() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("job 1 done.")
		wg.Done()
	}()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("job 2 done.")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("All Done.")
}

func stop2() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("got the stop channel")
				return
			default:
				fmt.Println("still working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("stop the gorutine")
	stop <- true
	time.Sleep(3 * time.Second)
}

func stop3() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("got the stop channel")
				return
			default:
				fmt.Println("still working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("stop the gorutine")
	cancel()
	time.Sleep(3 * time.Second)
}

func stop4() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func(ctx context.Context, duration time.Duration) {
		select {
		case <-ctx.Done():
			fmt.Println("handle:", ctx.Err())
		case <-time.After(duration):
			fmt.Println("process request with:", duration)
		}
	}(ctx, 500*time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("main:", ctx.Err())
	}
}
