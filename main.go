package main

import (
	"rpg/game"
	"rpg/utils"
)

func main() {
	utils.WriteLn("Make a choice to continue")
	utils.WriteLn("1. Start")

	ch, err := utils.ReadChar()
	if err != nil {
		panic(err)
	}

	if ch == '1' {
		game.Start()
	}
}
