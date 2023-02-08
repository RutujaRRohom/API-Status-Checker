package server 

import (
	"net/http"

	"github.com/RohomRutuja/Go_API/CheckStatus"
)


func InitRouter(dp *dependencies) {
	http.HandleFunc("/POST/websites", CheckStatus.PostHandler(dp.httpchecker))
	http.HandleFunc("/GET/websites", CheckStatus.GetHandler(dp.httpchecker))
}
