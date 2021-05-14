package router

import (
	"encoding/json"
	"net/http"
	"poe-debugserver/src/model"
	"strconv"
)

// Handler implement
func GetCharacterHandler(w http.ResponseWriter, r *http.Request) {
	//TODO:	to get character struct by character name and account name
}
func GetCharacterHandlerRandom(w http.ResponseWriter, r *http.Request) {
	//TODO:	to get character struct randomly, also give illegal struct
}
func GetLeagueListCurrent(w http.ResponseWriter, r *http.Request) {
	//TODO:	to get League info
}
func GetLeagueListRandom(w http.ResponseWriter, r *http.Request) {
	//TODO:	to get League info struct randomly, also give illegal struct
}
func GetItems(w http.ResponseWriter, r *http.Request) {
	//TODO:	to get character items struct.
}
func GetPassiveTree(w http.ResponseWriter, r *http.Request) {
	//TODO:	to get character's Passivetree struct.
}
func GetLadderHandler(w http.ResponseWriter, r *http.Request) {
	//TODO:	to get Ladder of league
}

func GetAccountNameHandler(w http.ResponseWriter, r *http.Request) {
	//TODO:	to get accountname struct (*but only one result)
	result := model.AccountName{
		Accountname: `defaultAccount`,
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		errorstruct := model.ErrorResult{
			Err: struct {
				Code    int8   "json:\"code\""
				Message string "json:\"message\""
			}{
				Code:    -99,
				Message: `json Marshal wrong. (from Debug Server)`,
			},
		}
		errorbytes, _ := json.Marshal(errorstruct)
		w.Write(errorbytes)
		return
	}
	w.Write(bytes)
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	qurey := r.URL.Query()
	mBs := qurey.Get("maintenane")
	mB, _ := strconv.ParseBool(mBs)
	Maintenance = mB
}
