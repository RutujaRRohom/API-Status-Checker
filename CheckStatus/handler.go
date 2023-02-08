package CheckStatus

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func PostHandler(httpcheck WebsiteChecker) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		if req.Method == "POST" {
			websiteMap := make(map[string][]string)
			json.NewDecoder(req.Body).Decode(&websiteMap)
			websitelist := websiteMap["websites"]
			for _, v := range websitelist {
				Webmap[v] = Website{v, "..."}
			}
			fmt.Println("post success")
			w.WriteHeader(http.StatusOK)

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

}

func GetHandler(httpcheck WebsiteChecker) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			//defer r.Body.Close()
			statusmap := make(map[string]string)
			web := r.URL.Query().Get("url")

			if web != "" {
				status, _ := httpcheck.CreateStatus(context.Background(), web)
				statusmap[web] = status
			} else {
				for URL := range Webmap {
					statusmap[URL], _ = httpcheck.CreateStatus(context.Background(), URL)
				}
			}
			json.NewEncoder(w).Encode(statusmap)

		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	})
}
