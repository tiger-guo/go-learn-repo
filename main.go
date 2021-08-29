/**
* @program: go-learn-repo
* @description:
* @author: guohuliu
* @create: 2021-08-29 21:56
**/

package main

import "go-learn-repo/sync"

func main() {
	/*
		Note: 不加锁，sum的值将会随机（sum <= 1000），加锁可以保证数据并发安全(sum == 1000)。
	*/
	sync.Mutex()

	/*
		Note：协程运行时间不确定的话，如果主程序结束，则会导致协程被动结束，WaitGroup则可以用于
		判断协程运行是否结束。
	*/
	sync.WaitGroup()

	/*
		Note: sync.Once可以确保代码只执行一次。
	*/
	sync.Once()

	/*
		Note: sync.Cond 用于需要等待所有协程准备就绪的情况下，用 cond 去唤醒所有协程去同时运行。
			比如，赛跑时需要等待所有人准备就绪，才可以比赛开始。
	*/
	sync.Cond()
}
