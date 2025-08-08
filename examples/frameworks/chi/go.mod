module chi-example

go 1.22

toolchain go1.23.6

require (
	github.com/go-chi/chi/v5 v5.2.2
	go.rumenx.com/feed v1.0.0
	go.rumenx.com/feed/adapters/chi v0.0.0-20250801144943-e24fced1544a
)

replace go.rumenx.com/feed => ../../..

replace go.rumenx.com/feed/adapters/chi => ../../../adapters/chi
