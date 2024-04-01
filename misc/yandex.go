package main

import "fmt"

func longOnes() {
	var n int
	fmt.Scanf("%d\n", &n)
	var longSoFar int
	var res int
	for i := 0; i < n; i++ {
		var e int
		fmt.Scanf("%d\n", &e)
		if e == 1 {
			longSoFar++
		} else {
			if longSoFar > res {
				res = longSoFar
			}
			longSoFar = 0
		}
	}
	if res < longSoFar {
		res = longSoFar
	}
	fmt.Println(res)
}

func removeDups() {
	var n int
	fmt.Scanf("%d\n", &n)
	seen := make(map[int]struct{})
	for i := 0; i < n; i++ {
		var e int
		fmt.Scanf("%d\n", &e)
		if _, in := seen[e]; !in {
			seen[e] = struct{}{}
			fmt.Println(e)
		}
	}
}

func isAnagram() {
	var f, s string
	fmt.Scanf("%s\n", &f)
	fmt.Scanf("%s\n", &s)
	if len(f) != len(s) {
		fmt.Println(0)
	}
	c := 0
	for i := 0; i < len(f); i++ {
		for j := 0; j < len(s); j++ {
			if f[i] == s[j] {
				c++
			}
		}
	}
	if len(f) == c {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

func parGen() {
	var n int
	fmt.Scanf("%d\n", &n)
	var res []string

	backtrack(n, 0, 0, "", &res)

	for i := range res {
		fmt.Println(res[i])
	}

}

func backtrack(n, o, c int, s string, r *[]string) {
	if o == n && c == n {
		*r = append(*r, s)
		return
	}

	if o < n {
		backtrack(n, o+1, c, s+"(", r)
	}

	if o > c {
		backtrack(n, o, c+1, s+")", r)
	}
}
