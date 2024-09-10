module golang-crawler

go 1.22.5

replace dataworker => ./pkg/dataworker

replace loader => ./pkg/loader

replace crawler => ./pkg/crawler

require (
	dataworker v0.0.0-00010101000000-000000000000
	loader v0.0.0-00010101000000-000000000000
)

require (
	crawler v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.29.0 // indirect
)
