package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(3)

// 	chA := make(chan bool)
// 	chB := make(chan bool)
// 	chC := make(chan bool)

// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 10; i++ {
// 			<- chA
// 			fmt.Println("goroutine 1 prints: a")
//             chB <- true
// 		}
//         <- chA
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 10; i++ {
// 			<- chB
// 			fmt.Println("goroutine 2 prints: b")
//             chC <- true
// 		}
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 10; i++ {
// 			<- chC
// 			fmt.Println("goroutine 3 prints: c")
//             chA <-  true
// 		}
// 	}()

//     chA <- true
//     wg.Wait()
//     close(chA)
//     close(chB)
//     close(chC)
// }



func main() {

	var wg sync.WaitGroup
	wg.Add(3)
	chA := make(chan bool)
	chB := make(chan bool)
	chC := make(chan bool)

	go func(s string){
		defer wg.Done()
		for i:=0; i<5; i++ {
		<- chA 
		fmt.Println(s)
		chB <- true
		}

		<- chA
	}("A")

	go func(s string){
		defer wg.Done()
		for i:=0; i<5; i++ {
		<- chB 
		fmt.Println(s)
		chC <- true
		}
	}("B")

	go func(s string){
		defer wg.Done()
		for i:=0; i<5; i++ {
		<- chC 
		fmt.Println(s)
		chA <- true
		}
	}("C")

	chA <- true
	wg.Wait()
	close(chA)
	close(chB)
	close(chC)

}