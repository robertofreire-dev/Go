# Go

## first program

go run app.go

Go compiles your program to a temporary directory, producing a binary, It immediately runs that binary and after execution, the binary and build artifacts are deleted automatically.

You can see the actual temp directory by using: go build -work app.go

## packages and module

to create a module: go mod init packagesandmoduleexample

this folder where go.mod lives and all subfolders belongs to module packagesandmoduleexample

to use all Go files in the folder, use: go run . 