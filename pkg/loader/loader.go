package loader

import (
	"fmt"
	"time"
)

type Loader struct {
	stop chan bool
}

func NewLoader() *Loader {
	return &Loader{
		stop: make(chan bool),
	}
}

func (l *Loader) Start() {
	// Символы для лоадера
	loaderChars := `|/-\`

	go func() {
		i := 0
		for {
			select {
			case <-l.stop:
				fmt.Print("\r")
				return
			default:
				fmt.Printf("\r%c", loaderChars[i%len(loaderChars)])
				i++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}

func (l *Loader) Stop() {
	l.stop <- true
}
