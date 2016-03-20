# fstord (WIP)
===

fstord (read [First Order](https://en.wikipedia.org/wiki/First_Order_(Star_Wars)) - yeah you're right) is a package that implemented [higher order functions](https://en.wikipedia.org/wiki/Higher-order_function).

## Install

`go get github.com/jaxi/fstord`

## Examples

```go


slice := []int{1, 2, 3, 4, 5}

fstord.Map(slice, func(x int) int { return x*x }).([]int)
# => [1, 4, 9, 16, 25]

fstord.Filter(slice, func(x int) bool { return x%2==1 }).([]int)
# => [1, 3, 5]

mp := map[int]int{
  1: 2,
  3: 4,
  5: 6,
}

fstord.Map(mp, func(a, b int) string {
  return strconv.Itoa(b + 1)
}).(map[int]string)
# => {1 : "3", 3: "5", 5: "7"}

fstord.Filter(mp, func(a b int) bool { return a+b==3})
# => {1: 2}

fstord.Reduce([]int{1, 2, 3, 4}, func(acc, i int) int { return acc+i }, 0)
# => 10
```

## TODO

- [x] Map
- [x] Filter
- [ ] Reduce (map data structure support hasn't been done yet)
- [ ] Any
- [ ] Every
- [ ] Count
- [ ] FilterMap

And more? Just comment!
