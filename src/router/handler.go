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
	//Default character list
	result := []model.Character{
		{
			Name:            `Player01FromDBServer`,
			League:          `Standard`,
			ClassId:         0,
			AscendancyClass: 0,
			Class:           `Scion`,
			Level:           100,
			Experience:      4250334444,
			LastActive:      true,
		},
		{
			Name:            `Player02FromDBServer`,
			League:          `Standard`,
			ClassId:         0,
			AscendancyClass: 0,
			Class:           `Scion`,
			Level:           100,
			Experience:      4250334444,
			LastActive:      true,
		},
		{
			Name:            `Player03FromDBServer`,
			League:          `Standard`,
			ClassId:         0,
			AscendancyClass: 0,
			Class:           `Scion`,
			Level:           100,
			Experience:      4250334444,
			LastActive:      true,
		},
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		errorstruct := model.ErrorResult{
			Err: struct {
				Code    int8   `json:"code"`
				Message string `json:"message"`
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

func GetLeagueListCurrent(w http.ResponseWriter, r *http.Request) {
	//TODO:	to get League info
	result := currentLeague
	bytes, err := json.Marshal(result)
	if err != nil {
		errorstruct := model.ErrorResult{
			Err: struct {
				Code    int8   `json:"code"`
				Message string `json:"message"`
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
				Code    int8   `json:"code"`
				Message string `json:"message"`
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
	//TODO: to set debug server status by GET method
	/*
		KEY:      			VALUE:
		maintenance		 	bool


	*/
	qurey := r.URL.Query()

	//Set Maintenance status
	mBs := qurey.Get("maintenance")
	mB, _ := strconv.ParseBool(mBs)
	Maintenance = mB
}
