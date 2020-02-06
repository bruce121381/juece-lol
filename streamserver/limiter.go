package main

import "log"

/**
用 bucket token 算法实现控流
*/
type ConnLimiter struct {
	/**
	bucket 容量 ---> 大小
	*/
	concurrentConn int

	bucket chan int
}

/**
构造函数
*/
func NewConnLimiter(concurrentConn int) *ConnLimiter {
	return &ConnLimiter{concurrentConn: concurrentConn, bucket: make(chan int, concurrentConn)}
}

//GetConn 获取token
func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reach the rate limitation")
		return false
	}

	cl.bucket <- 1

	return true

}

//RelaseConn 释放token
func (cl *ConnLimiter) ReleaseConn() {

	c := <-cl.bucket

	log.Printf("New Connection coming : %d", c)

}
