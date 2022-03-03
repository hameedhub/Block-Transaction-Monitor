package model

type RawBlock struct {
	Number string `json:"number"`
	Hash string `json:"hash"`
	Transaction []string `json:"transactions"`
}
