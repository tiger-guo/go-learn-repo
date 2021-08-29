/**
* @program: go-learn-repo
* @description:
* @author: guohuliu
* @create: 2021-08-29 22:13
**/

package sync

import (
	"fmt"
	"sync"
)

func Once() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	// 用于等待协程执行完毕
	var wg sync.WaitGroup
	wg.Add(10)
	// 启动10个协程执行once.Do(onceBody)
	for i := 0; i < 10; i++ {
		go func() {
			// 把要执行的函数(方法)作为参数传给once.Do方法即可
			once.Do(onceBody)
			defer wg.Done()
		}()
	}

	// wait
	wg.Wait()
}
