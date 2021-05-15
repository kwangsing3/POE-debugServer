package model

type Character struct {
	Name            string  `json:"name"`
	League          string  `json:"league"`
	ClassId         byte    `json:"classId"`
	AscendancyClass int8    `json:"ascendancyclass"`
	Class           string  `json:"class"`
	Level           byte    `json:"level"`
	Experience      float64 `json:"experience"`
	LastActive      bool    `json:"lastActive"`
}
