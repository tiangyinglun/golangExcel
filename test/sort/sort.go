package main

import (
	"fmt"
	"sort"
)

type MapSorter []Item

type Item struct {
	Key string
	Val int64
}

func NewMapSorter(m map[string]int64) MapSorter {
	ms := make(MapSorter, 0, len(m))

	for k, v := range m {
		ms = append(ms, Item{k, v})
	}

	return ms
}

func (ms MapSorter) Len() int {
	return len(ms)
}

func (ms MapSorter) Less(i, j int) bool {
	return ms[i].Val < ms[j].Val // 按值排序
	//return ms[i].Key < ms[j].Key // 按键排序
}

func (ms MapSorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func main() {
	m := map[string]int64{
		"e": 10,
		"a": 2,
		"d": 15,
		"c": 8,
		"f": 1,
		"b": 12,
	}

	ms := NewMapSorter(m)
	sort.Sort(ms)

	for _, item := range ms {
		fmt.Printf("%s:%d\n", item.Key, item.Val)
	}
}
