package game

import (
	"fmt"
	"rpg/game/story"
	"rpg/utils"
	"strconv"
)

func Start() {
	storyManager, err := story.NewStoryManager("resources/story.yaml")
	if err != nil {
		utils.WriteLn(fmt.Sprintf("Error loading story: %v", err))
		return
	}

	for !storyManager.IsEnd() {
		utils.ClearScreen()
		utils.WriteLn(storyManager.CurrentText())
		utils.WriteLn("") // Empty line for spacing

		options := storyManager.CurrentOptions()
		for i, option := range options {
			utils.WriteLn(fmt.Sprintf("%d. %s", i+1, option.Text))
		}

		utils.WriteText("\nEnter your choice: ")
		ch, err := utils.ReadChar()
		if err != nil {
			utils.WriteLn(fmt.Sprintf("Error reading input: %v", err))
			return
		}

		choice, err := strconv.Atoi(string(ch))
		if err != nil || choice < 1 || choice > len(options) {
			utils.WriteLn("\nInvalid choice. Press any key to try again.")
			utils.ReadChar() // Wait for any key press
			continue
		}

		storyManager.ChooseOption(choice - 1)
	}

	utils.ClearScreen()
	utils.WriteLn("The end of the story.")
}
