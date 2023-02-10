package main

import (
	"fmt"
	"strconv"

	"github.com/RohomRutuja/Go_API/CheckStatus"
	"github.com/RohomRutuja/Go_API/server"
	"github.com/urfave/negroni"
)

func main() {
	dep := server.InitDependencies()
	router := server.InitRouter(dep)
	go CheckStatus.GetStatus()

	server := negroni.Classic()
	server.UseHandler(router)
	port := 8080 // This can be changed to the service port number via environment variable.
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))
	server.Run(addr)

}
