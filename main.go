package main

import (
	"dataworker"
	"flag"
	"fmt"
	"loader"
)

func main() {
	urls := []string{"https://go.dev/", "https://golang.org"}

	l := loader.NewLoader()

	name := flag.String("name", "default", "name")
	flag.Parse()
	fmt.Printf("Parsing by this value: %v\n", *name)

	l.Start()

	slice := dataworker.GetDataAndSort(urls, *name)

	l.Stop()

	if len(slice) == 0 {
		fmt.Println("No matches found")
		return
	}

	for i, v := range slice {
		fmt.Printf("%v: %v - '%v'\n", i+1, v.Url, v.Name)
	}
}
