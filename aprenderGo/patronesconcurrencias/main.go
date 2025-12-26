package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fetchURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n ", url, err)
		return
	}

	defer resp.Body.Close()
	duracion := time.Since(start)

	fmt.Printf("fetched %s en %v - Status %x\n", url, duracion, resp.Status)
}

func main() {

	urls := []string{
		"https://rickandmortyapi.com/api/character/1",
		"https://rickandmortyapi.com/api/character/2",

		"https://rickandmortyapi.com/api/character/3",
		"https://rickandmortyapi.com/api/character/4",
	}

	start := time.Now()

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go fetchURL(url, &wg)
	}

	wg.Wait()

	duracionTotal := time.Since(start)
	fmt.Printf("Tiempo total concurrente: %v \n", duracionTotal)
}
