package types

type Response struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

type LoginInfos struct {
	Success bool `json:"success"`
}
