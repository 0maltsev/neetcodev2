package main

import(
	"sort"
)

type TimeMap struct {
	m map[string][]pair
}

type pair struct {
	timestamp int
	value string
}

func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]pair),
	}
}

 func (tm *TimeMap) Set(key string, value string, timestamp int)  {
	tm.m[key] = append(tm.m[key], pair{timestamp, value})
}

 func (tm *TimeMap) Get(key string, timestamp int) string {
	if _, exists := tm.m[key]; !exists {
		return ""
	}
	
	pairs := tm.m[key]
	idx := sort.Search(len(pairs), func(i int) bool {
		return pairs[i].timestamp > timestamp
	})
	
	if idx == 0 {
		return ""
	}
	return pairs[idx-1].value
}