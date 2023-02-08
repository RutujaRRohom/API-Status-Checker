package CheckStatus

import (
	"context"
	"errors"
	"net/http"
	"time"
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
		for k := range Webmap {
			resp, err := http.Get("http://" + k)
			flag := Webmap[k]
			if resp.StatusCode != http.StatusOK || err != nil {
				flag.Status = "DOWN"
				Webmap[k] = flag
			} else {
				flag.Status = "UP"
				Webmap[k] = flag
			}
		}
		time.Sleep(1 * time.Minute)
	}
}
