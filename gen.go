package main

//go:generate wget https://github.com/mantisbt/mantisbt/raw/master/api/rest/swagger.json
//go:generate go get github.com/go-swagger/go-swagger/cmd/swagger
//go:generate swagger generate client