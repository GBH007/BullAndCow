package main

func main() {
	v := NewConsoleView(NewGamesController())
	v.Run()
}
