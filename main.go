package main

import (
	"net/http"

	"github.com/RohomRutuja/Go_API/CheckStatus"
	"github.com/RohomRutuja/Go_API/server"
)


func main() {
	dep := server.InitDependencies()
	server.InitRouter(dep)
	go CheckStatus.GetStatus()
	http.ListenAndServe(":8080", nil)
}
