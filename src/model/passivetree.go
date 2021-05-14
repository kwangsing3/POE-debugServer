package model

type PassiveTree struct {
	Hashes           []int32  `json:"hashes"`
	Items            []Items  `json:"items"`
	Visual_overrides []string `json:"visual_overrides"`
}
