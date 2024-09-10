module golang-crawler

go 1.22.5

require (
	crawler v0.0.0-00010101000000-000000000000
	loader v0.0.0-00010101000000-000000000000
)

require golang.org/x/net v0.29.0 // indirect

replace loader => ./pkg/loader

replace crawler => ./pkg/crawler
