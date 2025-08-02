module chi-example

go 1.22

toolchain go1.23.6

require (
	github.com/go-chi/chi/v5 v5.0.11
	github.com/rumendamyanov/go-feed v1.0.0
	github.com/rumendamyanov/go-feed/adapters/chi v0.0.0-20250801144943-e24fced1544a
)

replace github.com/rumendamyanov/go-feed => ../../..
