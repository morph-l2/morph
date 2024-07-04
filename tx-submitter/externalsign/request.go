package externalsign

type SignData struct {
	Sign string `json:"sign"`
}

type Result struct {
	SignDatas []SignData `json:"signDatas"`
	Address   string     `json:"address"`
	SignType  int        `json:"signType"`
}

type Response struct {
	Result Result `json:"result"`
	Status int    `json:"status"`
}
