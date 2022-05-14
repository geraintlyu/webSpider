package main

import (
	"Spider/Spider/src/spider"
	"sync"
)

func main() {
	var wg = new(sync.WaitGroup)

	func() {
		wg.Add(1)
		spider.Zhihu(wg)
	}()

	func() {
		wg.Add(1)
		go spider.Bilibili(wg)
	}()
	func() {
		wg.Add(1)
		go spider.Nintendo(wg)
	}()

	wg.Wait()
}
