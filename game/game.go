package game

import (
	"fmt"
	"log"
	"rpg/game/core/player"
)

func Game() {
	player, err := player.CreatePlayer()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(player)
}
