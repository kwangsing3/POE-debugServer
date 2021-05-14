package model

type Character struct {
	AccountName     string `json:"accountname"`
	CharacterName   string `json:"charactername"`
	Server          string `json:"server"`
	Language        string `json:"language"`
	League          string `json:"league"`
	Level           byte   `json:"level"`
	Class           string `json:"class"`
	Ascendancy      string `json:"ascendancy"`
	AscendancyClass int8   `json:"ascendancyclass"`
	Depth           int    `json:"depth"`
	SingleDepth     int    `json:"singledepth"`
}
