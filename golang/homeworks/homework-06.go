package main

import (
	"fmt"
	"sync"
)

// for Part A
func say(text string, wg *sync.WaitGroup) {
	fmt.Printf("Its %s\n", text)
	wg.Done()
}

func partA() {
	texts := [3]string{"Hello", "World", "Go"}
	var wg sync.WaitGroup

	for _, text := range texts {
		wg.Add(1)
		go say(text, &wg)
	}
	wg.Wait()
}

func fibonacci(n int, c chan int) { // 호출할 횟수 n과 int type의 채널 c를 받는다
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x // x를 c에 송신한다
		x, y = y, x+y
	}
	close(c) // 10회의 반복이 끝났다면 c 채널을 닫는다
}

func partB() {
	c := make(chan int, 10) // 10의 버퍼 크기를 갖는 int type의 채널을 생성한다
	go fibonacci(cap(c), c) // fibonacci를 goroutine으로 호출, cap(c)는 10이며(버퍼 크기) c는 채널이다
	for i := range c {      // c 채널이 닫힐때까지 반복한다
		fmt.Println(i)
	}
}

func main() {
	partA()
	partB()
}
