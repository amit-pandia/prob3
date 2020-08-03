package main

import (
	"fmt"
)

type test struct {
	n      int
	a      []int
	b      []int
	c      []int
	result int
}

func main() {

	t := test{}

	_, err := fmt.Scanf("%d", &t.n)
	if err != nil {
		return
	}
	if t.n < 1 && t.n > 40000 {
		return
	}

	t.a = make([]int, t.n)
	for i := 0; i < t.n; i++ {
		fmt.Scan(&t.a[i])
		if t.a[i] < 1 && t.a[i] > 1000000000 {
			return
		}
	}

	t.b = make([]int, t.n)
	for i := 0; i < t.n; i++ {
		fmt.Scan(&t.b[i])
		if t.a[i] < 1 && t.a[i] > 1000000000 {
			return
		}
	}

	t.c = make([]int, t.n)
	for i := 0; i < t.n; i++ {
		fmt.Scan(&t.c[i])
		if t.a[i] < 1 && t.a[i] > 1000000000 {
			return
		}
	}

	for i := 0; i < t.n; i++ {
		for j := 0; j < t.n; j++ {
			if t.a[i] < t.b[j] {
				for k := 0; k < t.n; k++ {
					if t.b[j] < t.c[k] {
						t.result++
					}
				}
			}
		}
	}

	fmt.Println("Output:")
	fmt.Println(t.result)
	return
}