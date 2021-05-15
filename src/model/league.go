package model

type League struct {
	ID          string `json:"id"`
	Realm       string `json:"realm"`
	Description string `json:"description"`
	Registerat  string `json:"registerAt"`
	URL         string `json:"url"`
	Startat     string `json:"startAt"`
	Endat       string `json:"endAt"`
	Delveevent  bool   `json:"delveEvent"`
	Rules       []Rule `json:"rules"`
}
type Rule struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
