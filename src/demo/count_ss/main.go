package main

import "fmt"

// 统计素数个数

func main() {

	i := violence(100)
	fmt.Println(i)
	i = countPrimes(100)
	fmt.Println(i)
}

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}

	isPrime := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if !isPrime[i] {
			for j := i * i; j < n; j += i {
				isPrime[j] = true // 将i的倍数标记来非素数
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if !isPrime[i] {
			count++
		}
	}
	return count
}

func violence(num int) int {
	count := 0

	for i := 2; i < num; i++ {
		if isPrime(i) {
			count++
		}
	}

	return count
}

func isPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
