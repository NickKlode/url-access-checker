package service

import (
	"sort"
	"time"
)

func FindMax(list map[string]time.Duration) (string, time.Duration) {
	var td []time.Duration

	for _, dur := range list {
		if dur > 0 {
			td = append(td, dur)
		}

	}

	sort.Slice(td, func(i, j int) bool {
		return td[i] < td[j]
	})

	max := td[len(td)-1]
	var url string
	for k, v := range list {
		if v == max {
			url = k
			break
		}
	}
	return url, max
}
