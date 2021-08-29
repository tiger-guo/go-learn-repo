/**
* @program: go-learn-repo
* @description:
* @author: guohuliu
* @create: 2021-08-29 22:08
**/

package sync

import (
	"fmt"
	"sync"
)

func WaitGroup() {
	var wg sync.WaitGroup
	//因为要监控110个协程，所以设置计数器为110
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			//计数器值减1
			defer wg.Done()
			add(10)
		}()
	}

	//一直等待，只要计数器值为0
	wg.Wait()

	fmt.Println("和为:", readSum())
}
