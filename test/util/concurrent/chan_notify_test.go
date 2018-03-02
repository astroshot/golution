package concurrent

import (
	"fmt"
	"testing"
	"time"
)

/*
例子中我们定义一个 stop 的 chan，通知他结束后台 goroutine。实现也非常简单，在后台 goroutine 中，使用 select 判断 stop 是否可以接收到值，如果可以接收到，就表示可以退出停止了；如果没有接收到，就会执行 default 里的监控逻辑，继续监控，只到收到 stop 的通知。

有了以上的逻辑，我们就可以在其他 goroutine 中，给 stop chan 发送值了，例子中是在 main goroutine 中发送的，控制让这个监控的 goroutine 结束。

发送了 stop <- true 结束的指令后，我这里使用 time.Sleep(5 * time.Second)故意停顿 5 秒来检测我们结束监控 goroutine 是否成功。如果成功的话，不会再有 `Monitor is on.` 的输出了；如果没有成功，监控 goroutine 就会继续打印 `Monitor is on.` 输出。
*/
func TestChanNotify(t *testing.T) {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Monitor is stoped.")
				return
			default:
				fmt.Println("Monitor is on.")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	fmt.Println("You can stop Monitor now.")
	stop <- true
	time.Sleep(5 * time.Second)
}
