package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fscan(in, &t)
	for ch := 0; ch < t; ch++ {
		var n, k int
		fmt.Fscan(in, &n, &k)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		diff := make([]int, n+1)
		cur_change := 0
		ok := true
		for i := 0; i < n; i++ {
			cur_change += diff[i]
			if a[i]-cur_change < 0 {
				ok = false
				break
			}
			if a[i]-cur_change > 0 {
				if i+k > n {
					ok = false
					break
				}
				delta := a[i] - cur_change
				cur_change += delta
				diff[i+k] -= delta
			}
		}
		if ok {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
