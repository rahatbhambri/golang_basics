package main

import "fmt"

func main() {
	l := []int{5, 2, 4, 3, 1}
	fmt.Println(l)

	// fmt.Println(bsort(l))
	fmt.Println(msort(l))

}

func bsort(l []int) []int {
	n := len(l)

	for i := 0; i < n; i += 1 {
		for j := i + 1; j < n; j += 1 {
			if l[i] > l[j] {
				l[i], l[j] = l[j], l[i]
			}
		}
	}

	return l
}

func msort(l []int) []int {

	fmt.Println("called")
	n := len(l)
	var l3 []int

	if n == 1 {
		return []int{l[0]}
	} else {
		l1 := msort(l[:n/2])
		l2 := msort(l[n/2:])

		a, b := len(l1), len(l2)
		i, j := 0, 0

		for {
			if i < a && j < b {
				if l1[i] < l2[j] {
					l3 = append(l3, l1[i])
					i += 1
				} else {
					l3 = append(l3, l2[j])
					j += 1
				}
			} else {
				break
			}
		}

		for {
			if i < a {
				l3 = append(l3, l1[i])
				i += 1
			} else {
				break
			}
		}

		for {
			if j < b {
				l3 = append(l3, l2[j])
				j += 1
			} else {
				break
			}
		}

	}

	return l3

}
