module golang-crawler

go 1.22.5

replace dataworker => ./pkg/dataworker

replace crawler => ./pkg/crawler

require dataworker v0.0.0-00010101000000-000000000000

require (
	crawler v0.0.0-00010101000000-000000000000 // indirect
	github.com/briandowns/spinner v1.23.1 // indirect
	github.com/fatih/color v1.7.0 // indirect
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/mattn/go-isatty v0.0.8 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/term v0.24.0 // indirect
)
