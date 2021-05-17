package router

// Handler implement
import (
	"encoding/json"
	"fmt"
	"net/http"
	"poe-debugserver/src/model"
	"strconv"
)

//GetCharacterHandler:	to get character struct by character name and account name
func GetCharacterHandler(w http.ResponseWriter, r *http.Request) {
	result := defaultCharactList
	bytes, err := json.Marshal(result)
	if err != nil {
		errbyte := returnErrorCode(-99)
		w.Write(errbyte)
		return
	}
	w.Write(bytes)
	return
}

//GetLeagueListCurrent:	to get League info
func GetLeagueListCurrent(w http.ResponseWriter, r *http.Request) {
	result := defaultLeague
	bytes, err := json.Marshal(result)
	if err != nil {
		errbyte := returnErrorCode(-99)
		w.Write(errbyte)
		return
	}
	w.Write(bytes)
}

var leaTURN = true //could be change type to enum or other variable type if need.
func GetLeagueListTurn(w http.ResponseWriter, r *http.Request) {
	result := map[bool]interface{}{true: defaultLeague, false: RitualLeague}[leaTURN]
	leaTURN = !leaTURN
	bytes, err := json.Marshal(result)
	if err != nil {
		errbyte := returnErrorCode(-99)
		w.Write(errbyte)
		return
	}
	w.Write(bytes)
}

//GetItems:	to get character items struct.
func GetItems(w http.ResponseWriter, r *http.Request) {
	qurey := r.URL.Query()
	acN := qurey.Get("accountName")
	if acN == "" {
		errbyte := returnErrorCode(-99)
		w.Write(errbyte)
		return
	}
	acN = `defaultAccount` /*Debug Server*/
	charName := qurey.Get("character")
	if charName == "" {
		errbyte := returnErrorCode(1)
		w.Write(errbyte)
		return
	}
	for i := range characterList {
		if characterList[i].Character.Name == charName {
			byt, _ := json.Marshal(characterList[i].Items)
			w.Write(byt)
			return
		}
	}
	//character not found
	errbyte := returnErrorCode(1)
	w.Write(errbyte)
}

/*
	GetPassiveTree: to get character's Passivetree struct.
	KEY:      			VALUE:
	accountName		 	string
	character			string
*/
func GetPassiveTree(w http.ResponseWriter, r *http.Request) {
	qurey := r.URL.Query()
	acN := qurey.Get("accountName")
	if acN == "" {
		errbyte := returnErrorCode(-99)
		w.Write(errbyte)
		return
	}
	acN = `defaultAccount` /*Debug Server*/
	charName := qurey.Get("character")
	if charName == "" {
		errbyte := returnErrorCode(1)
		w.Write(errbyte)
		return
	}
	for i := range characterList {
		if characterList[i].Character.Name == charName {
			byt, _ := json.Marshal(characterList[i].PassiveTree)
			w.Write(byt)
			return
		}
	}
	//character not found
	errbyte := returnErrorCode(1)
	w.Write(errbyte)
}

//GetLadderHandler: to get Ladder of league
/*
	GetPassiveTree: to get character's Passivetree struct.
	KEY:      			VALUE:
	limit		 		int (max:200)
	offset				int
*/
func GetLadderHandler(w http.ResponseWriter, r *http.Request) {
	qurey := r.URL.Query()
	//Set Maintenance status
	limitS := qurey.Get("limit")
	offsetS := qurey.Get("offset")
	limit, _ := strconv.ParseInt(limitS, 0, 16)
	offset, _ := strconv.ParseInt(offsetS, 0, 16)
	length := map[bool]int{true: 15000, false: len(defaultCharactList)}[len(defaultCharactList) > 15000]
	if limit < 1 || limit > 200 || int(offset+limit) > length {
		byt := returnErrorCode(2)
		w.Write(byt)
		return
	}
	if limit > int64(length) {
		limit = int64(length)
	}
	counter := 0
	i := 0
	var ent []model.Entry
	for range defaultCharactList {
		counter++
		if counter-1 < int(offset) {
			i++
			continue
		}
		ent = append(ent, model.Entry{
			Rank: i + 1,
			Dead: false,
			Character: struct {
				ID         string `json:"id"`
				Name       string `json:"name"`
				Level      int    `json:"level"`
				Class      string `json:"class"`
				Time       int    `json:"time"`
				Score      int    `json:"score"`
				Experience int64  `json:"experience"`
			}{
				ID:         ``,
				Name:       defaultCharactList[i].Name,
				Level:      int(defaultCharactList[i].Level),
				Class:      defaultCharactList[i].Class,
				Time:       0,
				Score:      3,
				Experience: int64(defaultCharactList[i].Experience),
			},
			Account: struct {
				Name       string `json:"name"`
				Realm      string `json:"realm"`
				Challenges struct {
					Total int `json:"total"`
				} `json:"challenges"`
			}{
				Name:  `defaultAccount`,
				Realm: `pc`,
				Challenges: struct {
					Total int `json:"total"`
				}{Total: 40},
			},
		})
		i++
		if counter > int(offset)+int(limit)-1 {
			break
		}
	}

	result := model.Ladder{
		Total:   length,
		Entries: ent,
	}
	byt, err := json.Marshal(result)
	if err != nil {
		eByt := returnErrorCode(-99)
		w.Write(eByt)
		return
	}
	w.Write(byt)
}

//GetAccountNameHandler: to get accountname struct (*but only one result)
func GetAccountNameHandler(w http.ResponseWriter, r *http.Request) {
	result := model.AccountName{
		Accountname: `defaultAccount`,
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		errbyte := returnErrorCode(-99)
		w.Write(errbyte)
		return
	}
	w.Write(bytes)
}

//StatusHandler: to set debug server status by GET method
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	/*
		KEY:      			VALUE:
		maintenance		 	bool
	*/
	qurey := r.URL.Query()
	//Set Maintenance status
	mBs := qurey.Get("maintenance")
	mB, _ := strconv.ParseBool(mBs)
	Maintenance = mB

	w.Write([]byte(`{"maintenance:` + fmt.Sprintln(Maintenance) + `"}`))
}

func returnErrorCode(code int8) (bytes []byte) {
	msg := ``
	switch code {
	case -99:
		msg = `Error from "Debug Server`
	case 1:
		msg = `Resource not found`
	case 2:
		msg = `Invalid query`
	case 3:
		msg = `Rate limit exceeded`
	case 4:
		msg = `Internal error`
	case 5:
		msg = `Unexpected content type`
	case 6:
		msg = `Unauthorized`
	case 7:
		msg = `Forbidden`
	case 8:
		msg = `Temporarily Unavailable`
	case 9:
		msg = `Method not allowed`
	case 10:
		msg = `Unprocessable Entity`
	}

	result := model.ErrorResult{
		Err: struct {
			Code    int8   `json:"code"`
			Message string `json:"message"`
		}{
			Code:    code,
			Message: msg,
		},
	}
	bytes, _ = json.Marshal(result)
	return
}
