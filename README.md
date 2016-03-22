# fstord
===

fstord (read [First Order](https://en.wikipedia.org/wiki/First_Order_(Star_Wars)) - yeah you're right) is a package that implemented [higher order functions](https://en.wikipedia.org/wiki/Higher-order_function).

## Install

`go get github.com/jaxi/fstord`

## Examples

```go

package main

import (
	"fmt"

	"github.com/jaxi/fstord"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(
		fstord.Map(slice, func(x int) int { return x * x }).([]int),
	)
	// [1 4 9 16 25]

	fmt.Println(
		fstord.Filter(slice, func(x int) bool { return x%2 == 1 }).([]int),
	)
	// [1 3 5]

	mp := map[int]int{
		1: 2,
		3: 4,
		5: 6,
	}

	fmt.Println(
		fstord.Map(mp, func(a, b int) int { return b + 1 }).(map[int]int),
	)
	// map[1:3 3:5 5:7]

	fmt.Println(
		fstord.Filter(mp, func(a, b int) bool { return a+b == 3 }),
	)
	// map[1:2]

	fmt.Println(
		fstord.Reduce([]int{1, 2, 3, 4}, func(acc, i int) int { return acc + i }, 0).(int),
	)
	// 10

	fmt.Println(
		fstord.Reduce(mp, func(a, k, v int) int { return a + v }, 0).(int),
	)
	// 12

	fmt.Println(
		fstord.Any([]int{1, 2, 3, 4}, func(i int) bool { return i > 4 }),
	)
	// false

	fmt.Println(
		fstord.Every([]int{1, 3, 5}, func(i int) bool { return i%2 == 1 }),
	)
	// true

	fmt.Println(
		fstord.Count([]int{1, 2, 3, 4, 5}, func(i int) bool { return i%2 == 1 }),
	)
	// 3
}
```
