package models

type Keycode struct {
	Code string `json:"code" param:"code"`
}

type KeyStroke struct {
	IsSpecial bool   `json:"isSpecial" param:"isSpecial"`
	Key       string `json:"key" param:"key"`
}
