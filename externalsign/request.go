package externalsign

type SignData struct {
	Sign string `json:"sign"`
}

type Result struct {
	Address  string `json:"address"`
	SignType int    `json:"signType"`
	Sha3     string `json:"sha3"`
}

type Response struct {
	Result Result `json:"result"`
	Status int    `json:"status"`
}
