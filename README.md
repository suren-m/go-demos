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

## Create separate files under same package to explore basic constructs

1. Data types

2. Loops and Conditionals

3. Standard library 

4. Logging

5. Structs

6. Pointers

### Showcase the lack of encapsulation despite having separate files. This will be covered using local packages in next demo