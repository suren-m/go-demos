# go-demos

## First program

1. Write a simple program in `main.go` 

2. Show `go run main.go` and `go build main.go && ./main`

3. Run `go run .` to show module init message.

## Modules

1. Init a module with `go mod init github.com/suren-m/go-demos`

2. Create a simple program in `main.go`

3. Do a `go run .`

4. Do a `go fmt`

5. Do a `go build . && ./go-demos` (notice the name of the module used for binary)

6. Move `main.go` into `app` dir and run from there.

## Use external Modules

1. import `rsc.io/quote` and call `quote.Go()`

2. To install the module, Run `go mod tidy` to update `go.mod` and include `go.sum`

> see: https://github.com/rsc/quote/tree/v4.0.1
> also: https://pkg.go.dev/rsc.io/quote/v4

3. Navigate to `GOPATH` to show the location of source code for quote in local system (`go/pkg/mod/rsc.io/quote@vx.x.x`)

## Debug the project in vs code

1. InmsSetup launch config

## Create separate files under same package to explore basic constructs

1. Data types

2. Loops and Conditionals

3. Standard library 

4. Logging

5. Structs

6. Pointers

### Showcase the lack of encapsulation despite having separate files. This will be covered using local packages in next demo