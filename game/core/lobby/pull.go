package lobby

import "rpg/game/core/character"

func Pull() (character.Character, error) {
	pulledCharacter, err := character.CreateCharacter()
	if err != nil {
		return character.Character{}, err
	}

	return pulledCharacter, nil
}
