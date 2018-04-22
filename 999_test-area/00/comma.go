package main

import "fmt"

func main() {
	fmt.Println(comma("12359"))      // 12,359
	fmt.Println(comma("123590"))     // 123,590
	fmt.Println(comma("1235900"))    // 1,235,900
	fmt.Println(comma("1235900000")) // 1,235,900,000
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	fmt.Printf(" COMMA STRING = %s\n s[n-3:] = %v\n s[:n-3] = %v\n [n - 3] = %v\n n = %v\n", s, s[n-3:], s[:n-3], n-3, n)
	return comma(s[:n-3]) + "," + s[n-3:]
}
