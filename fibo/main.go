package fibo

func FiboRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FiboRecursion(n-1) + FiboRecursion(n-2)
}

func FiboMemo(n int, dp map[int]int) int {
	if n <= 1 {
		return n
	}
	dp[n] = FiboMemo(n-1, dp) + FiboMemo(n-2, dp)
	return dp[n]
}

func FiboTab(n int) int {
	if n <= 1 {
		return n
	}
	a := make([]int, n+1)
	a[0], a[1] = 0, 1
	for i := 2; i <= int(n); i++ {
		a[i] = a[i-1] + a[i-2]
	}
	return a[n]
}

func FiboSpaceOptimization(n int) int {
	a, b := 1, 1
	for i := 2; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

// func main() {
// 	dp_fibo := map[int]int{1: 0, 2: 1}
// 	fmt.Println(fiboRecursion(5))
// 	fmt.Println(fiboMemo(5, dp_fibo))
// 	fmt.Println(fiboTab(5))
// 	fmt.Println(fiboSpaceOptimization(5))
// }
