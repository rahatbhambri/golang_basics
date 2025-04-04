package main

import "fmt"

func RunCompaction(hm map[int]int) {
	for k, _ := range hm {
		hm[k] = find(k, hm)
	}
}

func union(a int, b int, hm map[int]int) bool {
	par_a := find(a, hm)
	par_b := find(b, hm)

	if par_a == par_b {
		return false
	} else if par_a < par_b {
		hm[par_b] = par_a
	} else {
		hm[par_a] = par_b
	}

	RunCompaction(hm)
	return true
}

func find(a int, hm map[int]int) int {
	if hm[a] == a {
		return hm[a]
	}

	// fmt.Println(hm)

	hm[a] = find(hm[a], hm)
	return hm[a]
}

func main() {

	n := 9
	hm := make(map[int]int, n)

	for i := 0; i < n; i += 1 {
		hm[i] = i
	}

	union(1, 2, hm)
	fmt.Println(hm)
	union(4, 5, hm)
	fmt.Println(hm)
	union(3, 4, hm)
	fmt.Println(hm)
	union(1, 5, hm)
	fmt.Println(hm)

}
