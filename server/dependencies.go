package server

import (
	"github.com/RohomRutuja/Go_API/CheckStatus"
)
type dependencies struct {
	httpchecker CheckStatus.WebsiteChecker
}

// InitDependencies initializes the router with the dependencies


func InitDependencies() *dependencies {
	return &dependencies{
		httpchecker: CheckStatus.NewFunc(),
	}
}
