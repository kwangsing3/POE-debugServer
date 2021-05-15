package router

import (
	"net/http"
	"path"
	"poe-debugserver/src/model"

	"github.com/gorilla/mux"
)

var DataCenter *model.DataCenter
var Maintenance bool = false

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello debug World!"))
	path := map[bool]string{true: path.Join(`HTML/maintenance.html`), false: path.Join(`HTML/index.html`)}[Maintenance]

	http.ServeFile(w, r, path)
}

//SetRoute: setting router path, you can add whatever path you want.
func SetRoute(r *mux.Router) {
	// load sample
	sandbox()
	DataCenter = &model.DataCenter{}
	DataCenter.Init()

	//API document
	r.HandleFunc(`/`, HomeHandler).Methods("GET")
	//Status contral
	r.HandleFunc(`/status`, StatusHandler).Methods("GET")

	r.HandleFunc(`/character-window/get-account-name-by-character`, GetAccountNameHandler).Methods("GET") //Get account name
	r.HandleFunc(`/character-window/get-characters`, GetCharacterHandler).Methods("GET")                  //Get Character info
	r.HandleFunc(`/api/leagues`, GetLeagueListCurrent).Methods("GET")                                     //Get LeagueList info
	r.HandleFunc(`/api/ladders`, GetLadderHandler).Methods("GET")                                         //Get Ladder

	//r.HandleFunc(``, ).Methods("GET")
}
