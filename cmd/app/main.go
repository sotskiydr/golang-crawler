package main

import (
	"flag"
	"fmt"
	"github.com/briandowns/spinner"
	"golang-crawler/pkg/dataworker"
	"time"
)

func main() {
	urls := []string{"https://go.dev/", "https://golang.org"}

	name := flag.String("name", "default", "name")
	flag.Parse()
	fmt.Printf("Parsing by this value: %s%v%s\n", "\033[32m", *name, "\033[0m")

	l := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	l.Color("red")

	l.Start()
	time.Sleep(4 * time.Second)

	slice := dataworker.GetDataAndSort(urls, *name)

	l.Stop()

	if len(slice) == 0 {
		fmt.Println("\033[31mNo matches found\033[0m")
		return
	}

	for i, v := range slice {
		fmt.Printf("%v: %v - '%v'\n", i+1, v.Url, v.Name)
	}
}
