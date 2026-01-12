# simple-cmp

- toy Go library inspired by `go-cmp`
- provide easy way to compare slices
- uses dynamic programming to find a shortest path between two slices
  by changing, removing, or inserting characters
- I use it in tests to assert slices

# usage

```go
want := []int{1, 2, 3, 4, 5}
got  := []int{0, 3, 4, 5, 6}

diff := cmp.Diff(want, got)
```

... the `diff` is ...

```
+ 1
0 > 2
3 = 3
4 = 4
5 = 5
6 -
```

... which means to get `want` from `got` ...

- insert `1`
- replace `0` with `2`
- leave `3`, `4`, and `5`
- remove `6`

if there is no difference, the `diff` is an empty string
