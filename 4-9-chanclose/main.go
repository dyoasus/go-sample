package main

import "fmt"

func main() {
	dataChan := make(chan int, 5)

	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() {
		<-syncChan1

		for {
			// 接收方依然可以正常接收到管道中的值
			if elem, ok := <-dataChan; ok {
				fmt.Printf("Received: %d, ok = %v [receiver]\n", elem, ok)
			} else {
				break
			}
		}

		fmt.Println("Done. [receiver]")
		syncChan2 <- struct{}{}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			dataChan <- i
			fmt.Printf("Sent: %d [sender] \n", i)
		}
		// 先关闭了通道
		close(dataChan)
		syncChan1 <- struct{}{}

		fmt.Println("Done. [sender]")

		syncChan2 <- struct{}{}

	}()

	<-syncChan2
	<-syncChan2
}
