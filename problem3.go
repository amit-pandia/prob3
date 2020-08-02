package main

import (
	"fmt"
	"strconv"
)

type testWrapper struct {
	TestCount int
	testData  test
}

type test struct {
	n      int
	a      []int
	b      []int
	c      []int
	result int
}

func main() {

	tw := testWrapper{}

	_, err := fmt.Scanf("%d", &tw.testData.n)
	if err != nil {
		return
	}

	length := tw.testData.n
	tw.testData.a = make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scan(&tw.testData.a[i])
		if tw.testData.a[i] < 1 && tw.testData.a[i] > 1000000000 {
			return
		}
	}

	tw.testData.b = make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scan(&tw.testData.b[i])
		if tw.testData.a[i] < 1 && tw.testData.a[i] > 1000000000 {
			return
		}
	}

	tw.testData.c = make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scan(&tw.testData.c[i])
		if tw.testData.a[i] < 1 && tw.testData.a[i] > 1000000000 {
			return
		}
	}

	if tw.testData.n < 1 && tw.testData.n > 40000 {
		return
	}

	m := map[string]string{}
	for i := 0; i < tw.testData.n; i++ {
		for j := 0; j < tw.testData.n; j++ {
			for k := 0; k < tw.testData.n; k++ {
				if tw.testData.a[i] == tw.testData.b[j] {
					if tw.testData.a[i] == tw.testData.b[k] {
						tw.testData.result++
					}
				}
				s := strconv.Itoa(tw.testData.a[i])
				s += "," + strconv.Itoa(tw.testData.b[j])
				s += "," + strconv.Itoa(tw.testData.c[k])
				if (tw.testData.a[i] != tw.testData.b[j]) && (tw.testData.b[j] != tw.testData.c[k]) && (tw.testData.a[i] != tw.testData.c[k]) {
					m[s] = s
				}
			}
		}
	}

	if tw.testData.result == 0 {
		tw.testData.result = len(m)
	}

	if tw.testData.n == 1 {
		if tw.testData.a[0] == tw.testData.b[0] {
			if tw.testData.a[0] == tw.testData.b[0] {
				tw.testData.result = 0
			}
		}
	}

	fmt.Println("Output:")
	fmt.Println(tw.testData.result)
	/*	for k, _ := range m {
		fmt.Println(k)
	}*/

	return
}
