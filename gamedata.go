package main

import "strings"

const (
	Bull  = "B"
	Cow   = "C"
	Empty = "_"
)

type GameData struct {
	number string
	turns  []string
}

func NewGameData(number string) *GameData {
	return &GameData{number: number}
}

func (gd *GameData) NewTurn(number string) (string, bool) {
	res := ""
	solved := 0
	for i := 0; i < 4; i++ {
		if number[i] == gd.number[i] {
			res += Bull
			solved++
			continue
		}
		if strings.Contains(gd.number, number[i:i+1]) {
			res += Cow
			continue
		}
		res += Empty
	}
	if solved == 4 {
		return res, true
	}
	return res, false
}

func (gd *GameData) GetNumberOfTurns() int { return len(gd.turns) }
