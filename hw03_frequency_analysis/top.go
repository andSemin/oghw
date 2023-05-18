package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type KeyVal struct {
	word string
	cnt  int
}

func Top10(s string) []string {
	var r []string
	words := strings.Fields(s)

	cnt := make(map[string]int)
	for _, k := range words {
		cnt[k]++
	}

	kvs := make([]KeyVal, len(cnt))
	i := 0
	for k, v := range cnt {
		kvs[i] = KeyVal{k, v}
		i++
	}

	if len(cnt) < 1 {
		return r
	}

	sort.Slice(kvs, func(i, j int) bool {
		if kvs[i].cnt != kvs[j].cnt {
			return kvs[i].cnt > kvs[j].cnt
		}
		return kvs[i].word < kvs[j].word
	})

	l := min(10, len(cnt))
	for i := 0; i < l; i++ {
		r = append(r, kvs[i].word)
	}

	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
