package CheckStatus

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"log"
)

type WebsiteChecker interface {
	CreateStatus(ctx context.Context, url string) (status string, err error)
}

func NewFunc() WebsiteChecker {
	return &Website{}
}

func (web Website) CreateStatus(ctx context.Context, url string) (status string, err error) {
	if _, ok := Webmap[url]; !ok {
		return "Key absent.", errors.New("Website Not Found")
	}

	return Webmap[url].Status, nil
}

func GetStatus() {
	for {
		fmt.Println("Webmap", Webmap)
		for k := range Webmap {
			fmt.Println("checking status for: ", k)
			resp, err := http.Get("http://" + k)
			if err != nil {
				log.Println(err)
				continue
			}
			flag := Webmap[k]
			if resp.StatusCode != http.StatusOK || err != nil {
				flag.Status = "DOWN"
				Webmap[k] = flag
			} else {
				flag.Status = "UP"
				Webmap[k] = flag
			}
		}
		
	}
}
