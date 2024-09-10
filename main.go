package main

import (
	"crawler/pkg/crawler/spider"
	"flag"
	"fmt"
	"loader"
	"strings"
)

type Res struct {
	url  string
	name string
}

func main() {
	l := loader.NewLoader()
	name := flag.String("name", "default", "name")
	flag.Parse()
	fmt.Printf("Parsing by this value: %v\n", *name)

	l.Start()

	s := spider.New()
	res, err := s.Scan("https://go.dev/", 2)

	if err != nil {
		fmt.Println(err)
		return
	}

	var slice []Res

	for i, v := range res {
		_ = i
		if strings.Contains(v.URL, *name) {
			slice = append(slice, Res{
				url:  v.URL,
				name: v.Title,
			})
		}
		//fmt.Printf("%T, %#v \n", v, v)
	}
	l.Stop()

	if len(slice) == 0 {
		fmt.Println("No matches found")
		return
	}

	for i, v := range slice {
		fmt.Printf("%v: %v - '%v'\n", i+1, v.url, v.name)
	}
}
