/**
* @program: go-learn-repo
* @description:
* @author: guohuliu
* @create: 2021-08-29 21:55
**/

package sync

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum   int
	mutex sync.Mutex
)

func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}

func Mutex() {
	for i := 0; i < 100; i++ {
		go add(10)
	}

	time.Sleep(1 * time.Second)

	fmt.Println("和为:", readSum())
}

func readSum() int {
	b := sum
	return b
}
