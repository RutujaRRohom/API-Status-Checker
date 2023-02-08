package CheckStatus

type Website struct{
	URL string `json:"url"`
    Status string `json:"status"`
}

var Webmap=make(map[string]Website)