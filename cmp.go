package cmp

import (
	"fmt"
	"strings"
)

type Direction int

const (
	Insert Direction = iota
	Delete
	Replace
	Same
)

func nDirections(n int, d Direction) []Direction {
	ds := make([]Direction, n)
	for i := range ds {
		ds[i] = d
	}
	return ds
}

func shortestPath[T comparable](s1, s2 []T) []Direction {
	// .I
	// DR
	rows, cols := len(s1)+1, len(s2)+1

	dpStep := make([][]int, rows)
	dpPath := make([][][]Direction, rows)

	for i := range dpStep {
		dpStep[i] = make([]int, cols)
		dpPath[i] = make([][]Direction, cols)
	}

	for c := 0; c < cols; c++ {
		dpStep[0][c] = c
		dpPath[0][c] = nDirections(c, Insert)
	}

	for r := 0; r < rows; r++ {
		dpStep[r][0] = r
		dpPath[r][0] = nDirections(r, Delete)
	}

	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			if s1[r-1] == s2[c-1] {
				dpStep[r][c] = dpStep[r-1][c-1]
				dpPath[r][c] = append(dpPath[r-1][c-1], Same)
				continue
			}

			replace := dpStep[r-1][c-1] + 1
			delete := dpStep[r-1][c] + 1
			insert := dpStep[r][c-1] + 1

			minStep := min(replace, delete, insert)

			if replace == minStep {
				dpPath[r][c] = append(dpPath[r-1][c-1], Replace)
			} else if delete == minStep {
				dpPath[r][c] = append(dpPath[r-1][c], Delete)
			} else {
				dpPath[r][c] = append(dpPath[r][c-1], Insert)
			}

			dpStep[r][c] = minStep
		}
	}

	return dpPath[rows-1][cols-1]
}

func areEqual[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func Diff[T comparable](want, got []T) string {
	if areEqual(want, got) {
		return ""
	}

	lines := []string{}
	i1, i2 := 0, 0

	path := shortestPath(got, want)
	for _, d := range path {
		switch d {
		case Same:
			lines = append(lines, fmt.Sprintf("%v = %v", got[i1], want[i2]))
			i1++
			i2++
		case Replace:
			lines = append(lines, fmt.Sprintf("%v > %v", got[i1], want[i2]))
			i1++
			i2++
		case Delete:
			lines = append(lines, fmt.Sprintf("%v -", got[i1]))
			i1++
		case Insert:
			lines = append(lines, fmt.Sprintf("+ %v", want[i2]))
			i2++
		default:
			panic("unreachable: unknown direction")
		}
	}

	return strings.Join(lines, "\n")
}
