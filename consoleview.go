package main

import (
	"bufio"
	"fmt"
	"os"
)

type ConsoleView struct {
	gamesController *GamesController
	game            Game
	scanner         *bufio.Scanner
}

func NewConsoleView(gamesController *GamesController) *ConsoleView {
	return &ConsoleView{
		gamesController: gamesController,
		scanner:         bufio.NewScanner(os.Stdin),
	}
}

func (cv *ConsoleView) text() string {
	cv.scanner.Scan()
	return cv.scanner.Text()
}

func (cv *ConsoleView) Run() {
	for {
		if cv.game == nil && !cv.newGame() {
			return
		}
		if cv.makeTurn() && !cv.newGame() {
			return
		}
	}
}
func (cv *ConsoleView) newGame() bool {
	fmt.Println("Введите new для начала новой игры или нажмите Enter чтобы выйти")
	switch cv.text() {
	case "new":
		game, _ := cv.gamesController.NewGame()
		cv.game = game
		return true
	default:
		return false
	}
}
func (cv *ConsoleView) makeTurn() bool {
	fmt.Println("Ход номер", cv.game.GetNumberOfTurns()+1)
	fmt.Println("Введите четырехзначтное число")
	res, win := cv.game.NewTurn(cv.text())
	fmt.Println(res)
	return win
}
