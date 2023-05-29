package main

// go:generate wget -O swagger.json https://github.com/mantisbt/mantisbt/raw/master/api/rest/swagger.json
//go:generate go install github.com/go-swagger/go-swagger/cmd/swagger@latest
//go:generate swagger generate client
//go:generate go mod tidy
