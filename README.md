# fstord (WIP)
===

fstord (read [First Order](https://en.wikipedia.org/wiki/First_Order_(Star_Wars)) - yeah you're right) is a package that implemented [higher order functions](https://en.wikipedia.org/wiki/Higher-order_function). Currently support Map.

## Install

`go get github.com/jaxi/fstord`

## Examples

```go
slice := []int{1, 2, 3, 4, 5}

fstord.Map(slice, func(x int) int { return x * x }).([]int)
# => [1, 4, 9, 16, 25]

mp := map[int]int{
  1: 2,
  3: 4,
  5: 6,
}

fstord.Map(mp, func(a, b int) string {
  return strconv.Itoa(b + 1)
}).(map[int]string)

# => {1 : "3", 3: "5", 5: "7"}
```
