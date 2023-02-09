package CheckStatus

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"log"
)

func AddWebsitesHandler(httpcheck WebsiteChecker) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		websiteMap := make(map[string][]string)
		err := json.NewDecoder(req.Body).Decode(&websiteMap)
		if err != nil {
			log.Println("error in reading req body", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//validate req body
		if len(websiteMap) == 0 {
			log.Println("error website list is empty")
			w.WriteHeader(http.StatusBadRequest)
			return

		}
		websitelist := websiteMap["websites"]
		for _, v := range websitelist {
			Webmap[v] = Website{v, ""}
		}
		w.WriteHeader(http.StatusOK)
		
		fmt.Println("post success")
	})

}

func GetWebsitesHandler(httpcheck WebsiteChecker) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		statusmap := make(map[string]string)
		web := r.URL.Query().Get("url")
		if web != "" {
			stat, _ := httpcheck.CreateStatus(context.Background(), web)
			statusmap[web] = stat
		} else {
			for URL := range Webmap {
				statusmap[URL], _ = httpcheck.CreateStatus(context.Background(), URL)
			}
		}
		json.NewEncoder(w).Encode(statusmap)
	})
}
