package main

import (
	"fmt"
	"sync"
	"time"
)



func main(){
	start := time.Now()
	userName := fetchUser()
	// Before go routines time := 350millisecond
	// likes := fetchUserLikers(userName)
	// match := fetchUserMatch(userName)

	// fmt.Println("Likes: ", likes)
	// fmt.Println("match: ", match)

	// After go routines time := 100ms
	
	// creating channel
	respch := make(chan any, 2)// creating responce channel that accept any value returned by function

	// creating wait group
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go fetchUserLikers(userName, respch, wg)
	go fetchUserMatch(userName, respch, wg)

	wg.Wait()// block until 2 wg.Done()

	close(respch)

	for resp := range respch {
		fmt.Println("Response: ", resp)
	}
	fmt.Println("took: ", time.Since(start))
	

}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "om"
}

func fetchUserLikers(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)
	respch <- 11
	wg.Done()
}

func fetchUserMatch(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	respch <- "anna"
	wg.Done()
}