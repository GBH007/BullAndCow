package main

import (
	"math/rand"
	"strconv"
)

type Game interface {
	NewTurn(number string) (string, bool)
	GetNumberOfTurns() int
}

type GamesController struct {
	games []*GameData
}

func NewGamesController() *GamesController { return &GamesController{} }
func (gc *GamesController) getNewNumber() string {
	num := rand.Intn(9000) + 1000
	return strconv.Itoa(num)
}
func (gc *GamesController) NewGame() (Game, int) {
	id := len(gc.games)
	game := NewGameData(gc.getNewNumber())
	gc.games = append(gc.games, game)
	return game, id
}

func (gc *GamesController) GetGame(id int) (Game, bool) {
	if id < len(gc.games) {
		return gc.games[id], true
	}
	return nil, false
}
