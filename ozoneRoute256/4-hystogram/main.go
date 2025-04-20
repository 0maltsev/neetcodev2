package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  var in *bufio.Reader
  var out *bufio.Writer

  in = bufio.NewReader(os.Stdin)
  out = bufio.NewWriter(os.Stdout)
  defer out.Flush()

  var t, N, M int
  fmt.Fscan(in, &t)

  for test := 0; test < t; test++ {
    fmt.Fscan(in, &N, &M)

    counter := make(map[[100]int]int)
    arrIdxs := make(map[[100]int][]int)

    for n := 0; n < N; n++ {
      var arr [100]int
      for i := 0; i < M; i++ {
        fmt.Fscan(in, &arr[i])
      }
      minVal := arr[0]
      for i := 1; i < M; i++ {
        if arr[i] < minVal {
          minVal = arr[i]
        }
      }
      for i := 0; i < M; i++ {
        arr[i] -= minVal
      }
      counter[arr]++
      arrIdxs[arr] = append(arrIdxs[arr], n)
    }

    answer := 0
    // pairs := make(map[[2]int]bool)

    for arr, c := range counter {
      for k := 0; k < c; k++ {
        reverseArray := [100]int{}
        maxVal, minVal := arr[0], arr[0]
        for j := 1; j < M; j++ {
          if arr[j] > maxVal {
            maxVal = arr[j]
          }
          if arr[j] < minVal {
            minVal = arr[j]
          }
        }
        for j := 0; j < M; j++ {
          reverseArray[M-j-1] = maxVal + minVal - arr[j]
        }

        if _, ok := counter[reverseArray]; ok {
          // idxsInput := arrIdxs[arr]
          // idxsReversed := arrIdxs[reverseArray]
          //
          // for _, idx1 := range idxsInput {
          //   for _, idx2 := range idxsReversed {
          //     if idx1 < idx2 {
          //       pair := [2]int{idx1, idx2}
          //       if !pairs[pair] {
          //         answer++
          //         pairs[pair] = true
          //       }
          //     }
          //   }
          // }
          // fmt.Println(counter[reverseArray])
          answer += counter[reverseArray]
          if arr == reverseArray {
            answer--
          }
          counter[arr]--
        }
      }
    }
    fmt.Println(answer)

  }
}