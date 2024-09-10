package dataworker

import (
	"crawler"
	"crawler/spider"
	"strings"
	"sync"
)

type Res struct {
	Url  string
	Name string
}

func GetDataAndSort(urls []string, name string) []Res {
	s := spider.New()
	results := make(chan []crawler.Document, len(urls))
	errors := make(chan error, len(urls))
	var slice []Res

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			res, err := s.Scan(url, 2)
			if err != nil {
				errors <- err
			}
			results <- res
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
		close(errors)
	}()

	for res := range results {
		for i, v := range res {
			_ = i
			if strings.Contains(v.URL, name) {
				slice = append(slice, Res{
					Url:  v.URL,
					Name: v.Title,
				})
			}
		}
	}
	//fmt.Printf("%T, %#v \n", v, v)
	return slice
}
