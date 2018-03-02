package concurrent

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// testWaitGroup controls multi goroutine finish at the same time using WaitGroup
/*
一个很简单的例子，一定要例子中的2个goroutine同时做完，才算是完成，先做好的就要等着其他未完成的，所有的goroutine要都全部完成才可以。

这是一种控制并发的方式，这种尤其适用于，好多个goroutine协同做一件事情的时候，因为每个goroutine做的都是这件事情的一部分，只有全部的goroutine都完成，这件事情才算是完成，这是等待的方式。

在实际的业务种，我们可能会有这么一种场景：需要我们主动的通知某一个goroutine结束。比如我们开启一个后台goroutine一直做事情，比如监控，现在不需要了，就需要通知这个监控goroutine结束，不然它会一直跑，就泄漏了。
*/
func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1 finished!")
		wg.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2 finished!")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Ok, all is done.")
}
