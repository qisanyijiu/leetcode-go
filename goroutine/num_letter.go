package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func PrintNumOrLetter() {
	numberCh := make(chan bool)
	letterCh := make(chan bool)
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		defer wg.Done()
		for i := 1; i <= 10; i += 1 {
			<-numberCh // 等待信号
			fmt.Println(i)
			letterCh <- true // 通知打印字母
		}
		close(letterCh)
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		for i := 'A'; i <= 'J'; i += 1 {
			<-letterCh
			fmt.Printf("%c\n", i)
			if i < 'J' {
				numberCh <- true
			}
		}
		close(numberCh)
	}()
	numberCh <- true
	time.Sleep(100 * time.Millisecond)
}
