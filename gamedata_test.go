package main

import (
	"strings"
	"testing"
)

func TestGameData(t *testing.T) {
	gd := NewGameData("1234")
	ans, slv := gd.NewTurn("4321")
	t.Log(ans, slv)
	if slv || ans != strings.Repeat(Cow, 4) {
		t.Fail()
	}
	ans, slv = gd.NewTurn("5678")
	t.Log(ans, slv)
	if slv || ans != strings.Repeat(Empty, 4) {
		t.Fail()
	}
	ans, slv = gd.NewTurn("1234")
	t.Log(ans, slv)
	if !slv {
		t.Fail()
	}
}

func BenchmarkNewTurn(b *testing.B) {
	gd := NewGameData("1234")
	req := []string{"1234", "4321", "5678"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gd.NewTurn(req[i%3])
	}
}
