package service

import (
	"inhouseadtest/internal/storage"
	"log"
	"net/http"
	"time"
)

func request(list map[string]time.Duration) {

	for url := range list {
		go makeRequest(list, url)
	}

}

func makeRequest(list map[string]time.Duration, url string) {

	start := time.Now()
	resp, err := http.Get("http://" + url)
	if err != nil {
		log.Printf("cannot to make get request to url: %s", url)
		return
	}
	duration := time.Since(start)

	log.Printf("url: %s, response time: %s, status code: %d", url, duration, resp.StatusCode)

	if resp.StatusCode == 200 {

		list[url] = duration
	} else {
		list[url] = 0
	}
}

func RequestLoop() {
	list := storage.SitesStorage
	for {
		request(list)
		time.Sleep(time.Minute)
	}
}
