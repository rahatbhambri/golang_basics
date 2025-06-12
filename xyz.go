package main

import (
	"encoding/json"
	"fmt"
	http "net/http"
	"os"
	"sync"

	chi "github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	r.Get("/getdata", getHandler)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}

}

type sample struct {
	A int `json:"a"`
	B int `json:"b"`
	C int `json:"c"`
}

type UserRequest struct {
	NumThreads int `json:"numthreads"`
}

func getHandler(w http.ResponseWriter, req *http.Request) {

	f, err := os.ReadFile("./sample.json")
	if err != nil {
		fmt.Println("error reading file")
		return
	}

	u1 := &UserRequest{}
	json.NewDecoder(req.Body).Decode(u1)

	// n := u1.NumThreads

	var s1 sample
	// fmt.Println(f)
	err = json.Unmarshal(f, &s1)
	if err != nil {
		fmt.Println("error in unmarshal")
		return
	}

	results := make(chan int, 3)

	wg := sync.WaitGroup{}

	wg.Add(3)
	go Processor(s1.A, &wg, results)
	go Processor(s1.B, &wg, results)
	go Processor(s1.C, &wg, results)

	wg.Wait()

	fmt.Println(s1)

	for v := range results {
		fmt.Println(v)
	}

	var arr []byte
	arr = []byte("Hello world")
	w.Write(arr)
	w.Write(f)
}

func Processor(v int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()
	results <- v * (v)
}
