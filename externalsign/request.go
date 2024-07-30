package externalsign

type SignData struct {
	PubData string `json:"pubData"`
	Sign    string `json:"sign"`
}

type Result struct {
	Address   string      `json:"address"`
	SignType  int         `json:"signType"`
	SignDatas []*SignData `json:"signDatas"`
}

type Response struct {
	Result Result `json:"result"`
	Status int    `json:"status"`
}
