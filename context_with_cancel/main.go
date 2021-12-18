package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		err := action_1(ctx)

		if err != nil { // action_1 failed then we need to cancel all the follow-up actions
			cancel()
		}
	}()

	go func() {
		ctx2, cancel2 := context.WithCancel(ctx)

		err := action_2(ctx2)
		if err != nil {
			cancel2()
		}

	}()

	action3(ctx)
	time.Sleep(1700 * time.Millisecond)

}

func action_1(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	return errors.New("failed")
}

func action_2(ctx context.Context) error {
	select {
	case <-time.After(1500 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("cancel action 2")
	}
	return nil
}

func action3(ctx context.Context) {
	fmt.Println("do step 1")
	select {
	case <-ctx.Done():
		return
	default:
	}
	fmt.Println("do step 2")
	select {
	case <-ctx.Done():
		return
	default:
	}
	fmt.Println("do step 3")
}

func subAction3(ch chan struct{}) {
	time.Sleep(1700 * time.Millisecond)
	fmt.Println("subAction3 finish execution")
	ch <- struct{}{}
}

//output:
// cancel action 2
// cancel action 3
