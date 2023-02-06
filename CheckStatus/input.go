package CheckStatus

type Website struct {
	Url    string `json:"url"`
	Status string `json:"status"`
}


var Webmap = make(map[string]Website)
