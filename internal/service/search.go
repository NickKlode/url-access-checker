package service

import "time"

func FindSiteRespTime(site string, list map[string]time.Duration) (time.Duration, bool) {
	for url, dur := range list {
		if url == site {
			return dur, true
		}
	}
	return 0, false
}
