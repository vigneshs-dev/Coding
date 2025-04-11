package main

import (
    "fmt"
)

func main() {
    var n, t int
    fmt.Scan(&n, &t)

    var s string
    fmt.Scan(&s)

    queue := []rune(s)

    for time := 0; time < t; time++ {
        i := 0
        for i < n-1 {
            if queue[i] == 'B' && queue[i+1] == 'G' {
                queue[i], queue[i+1] = queue[i+1], queue[i]
                i += 1 // skip the next index to avoid re-swapping
            }
            i++
        }
    }

    fmt.Println(string(queue))
}
