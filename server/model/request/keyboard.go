package model

type KeycodeResquest struct {
	Code string `json:"code" param:"code"`
}

type KeyStrokeResquest struct {
	IsSpecial bool   `json:"isSpecial" param:"isSpecial"`
	Key       string `json:"key" param:"key"`
}
