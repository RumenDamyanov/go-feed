module echo-example

go 1.23.0

toolchain go1.23.6

require (
	github.com/labstack/echo/v4 v4.12.0
	go.rumenx.com/feed v1.0.0
	go.rumenx.com/feed/adapters/echo v0.0.0-20250801144943-e24fced1544a
)

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.40.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	golang.org/x/time v0.5.0 // indirect
)

replace go.rumenx.com/feed => ../../..

replace go.rumenx.com/feed/adapters/echo => ../../../adapters/echo
