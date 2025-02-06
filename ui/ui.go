package ui

import (
	"cubes/main/ai"
	"cubes/main/engine"
	"fmt"
)

func MainMenu() {
	for {
		a, err := selectOption("Welcome to Sav's 3D 4 in a row. Please select an option:", []string{"Play singleplayer", "Play multiplayer (local)", "View rules", "Exit"})
		if err != nil {
			clearTerminal()
			fmt.Println("Bye :)")
			return
		}
		switch a {
		case 0:
			singlePlayer()
		// TODO: add local multiplayer and rules
		case 3:
			return
		}
	}
}

func singlePlayer() {
	var state engine.GameState = engine.CreateEmpty()
	var finished bool
	var winner engine.FieldState
	for {
		finished, winner = state.GetWinner()
		if finished {
			break
		}
		move, err := selectMove(&state)
		if err != nil {
			clearTerminal()
			fmt.Print("Bye :)")
			return
		}
		state = *state.GetMovedClone(move)
		clearTerminal()
		printState(&state)
		finished, winner = state.GetWinner()
		if finished {
			break
		}
		move = ai.GetNextMove(&state, 4)
		state = *state.GetMovedClone(move)
		clearTerminal()
		printState(&state)
	}
	switch winner {
	case engine.White:
		fmt.Println("You won!")
	case engine.Black:
		fmt.Println("You lost.")
	default:
		fmt.Println("It was a draw.")
	}
	fmt.Scanln()
}
