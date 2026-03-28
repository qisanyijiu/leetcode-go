package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func ProducerAndConsumer() {
	var (
		producerCnt      = 2
		consumerCnt      = 3
		itemsPerProducer = 10
	)
	dataCh := make(chan int, itemsPerProducer*consumerCnt)
	var producerWG sync.WaitGroup
	var consumerWG sync.WaitGroup
	producerWG.Add(producerCnt)
	consumerWG.Add(consumerCnt)
	for i := 1; i <= producerCnt; i++ {
		go func(id int) {
			defer producerWG.Done()
			for num := 0; num < itemsPerProducer; num++ {
				data := id*100 + num
				dataCh <- data
			}
		}(i)
	}

	for i := 1; i <= consumerCnt; i++ {
		go func(id int) {
			defer consumerWG.Done()
			for num := range dataCh {
				fmt.Printf("consumer %d produced %d\n", id, num)
				time.Sleep(20 * time.Millisecond)
			}
		}(i)
	}

	go func() {
		producerWG.Wait()
		close(dataCh)
		fmt.Println("所有生产者完成，channel 已关闭")
	}()
	consumerWG.Wait()
	time.Sleep(time.Second)
}
