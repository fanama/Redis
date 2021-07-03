package types

type Response struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

type TestPost struct {
	Success string `json:"success"`
}

type IP struct {
	Origin string `json:"origin"`
}
