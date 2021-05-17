package model

type Ladder struct {
	Total   int     `json:"total"`
	Entries []Entry `json:"entries"`
}

type Entry struct {
	Rank      int  `json:"rank"`
	Dead      bool `json:"dead"`
	Character struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Level      int    `json:"level"`
		Class      string `json:"class"`
		Time       int    `json:"time"`
		Score      int    `json:"score"`
		Experience int64  `json:"experience"`
	} `json:"character"`
	Account struct {
		Name       string `json:"name"`
		Realm      string `json:"realm"`
		Challenges struct {
			Total int `json:"total"`
		} `json:"challenges"`
	} `json:"account"`
}
