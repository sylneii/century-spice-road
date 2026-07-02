package main

import (
	"encoding/json"
	"net/http"

	"github.com/sylneii/century-spice-road/game"
)

func homePageHandler(w http.ResponseWriter, req *http.Request) {

}

type CenturyServer struct {
	gameState *game.GameState
}

type NewGame struct {
	Username string `json:"username"`
}

func NewGameHandler(w http.ResponseWriter, req *http.Request) {
	newGameStruct := NewGame{}
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&newGameStruct)
	if newGameStruct.Username == "" {
		return
	}
}

func ActionHandler(w http.ResponseWriter, req *http.Request) {
}

func main() {

	// ruleEngine := NewRuleEngine()

	mux := http.NewServeMux()

	mux.Handle("GET /", http.FileServer(http.Dir("/web")))
	mux.HandleFunc("POST /game", NewGameHandler)
	mux.HandleFunc("POST /apply_action", ActionHandler)
}
