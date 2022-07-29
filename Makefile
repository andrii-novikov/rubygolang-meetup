build:
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/hello_world cmd/hello_world/main.go
	env GOOS=linux GOARCH=arm go build -o bin/linux/arm/hello_world cmd/hello_world/main.go
	env GOOS=windows GOARCH=amd64 go build -o bin/windows/amd64/hello_world.exe cmd/hello_world/main.go
	env GOOS=windows GOARCH=386 go build -o bin/windows/386/hello_world.exe cmd/hello_world/main.go
	env GOOS=darwin GOARCH=amd64 go build -o bin/darwin/amd64/hello_world cmd/hello_world/main.go
	env GOOS=darwin GOARCH=arm64 go build -o bin/darwin/arm64/hello_world cmd/hello_world/main.go