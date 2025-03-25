package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mx sync.Mutex

func say() {
	mx.Lock()
	defer wg.Done()
	fmt.Println("hello")
	time.Sleep(2 * time.Second)
	mx.Unlock()
}

func perform() {
	defer handleErrors()
	a, b := 5, 0
	fmt.Println(a / b)
}

func handleErrors() {
	recover()
	fmt.Println("Recovered")
}

func main() {

	for j := 0; j < 5; j += 1 {
		wg.Add(1)
		go say()
	}

	wg.Wait()

	perform()

	fmt.Println("after words this executes")

}
