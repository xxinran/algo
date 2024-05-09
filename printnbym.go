package main

import (
	"fmt"
	"math"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	m := 5
	n := 100
	fmt.Println(math.Ceil(float64(5)/float64(2)))
	perGoroutine :=  n / m
	//fmt.Println(perGoroutine)
	ch := make(chan int, m)

	for i := 0; i < m; i++ {
		wg.Add(1)
		start := i * perGoroutine
		end :=  start + perGoroutine - 1

		if i == m - 1 {
			end = n
		}
		go func(start int, end int, wg *sync.WaitGroup){
			fmt.Println("start = ", start, "end = ", end)
			defer wg.Done()
			//fmt.Println(<-ch)
			//<- ch
			ch_head := <- ch
			fmt.Println("head = ", ch_head)
			if ch_head == start {
				for i:= start; i <= end; i++ {
					fmt.Println(i)
				}
				fmt.Println("end+1", end+1)
				ch <- end + 1
			} else {
				ch <- ch_head
			}
		}(start, end, &wg)

	}
	ch <- 0
	wg.Wait()
	//close(ch)

}
