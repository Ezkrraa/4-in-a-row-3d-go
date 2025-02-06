package main

import (
	"cubes/main/ai"
	"cubes/main/engine"
	"fmt"
)

func main() {
	testing()
}

func testing() {
	state := engine.CreateEmpty()
	state.Board = [4][4][4]engine.FieldState{
		{
			{engine.White, engine.White, engine.White, engine.Empty},
			{engine.Black, engine.Empty, engine.Empty, engine.Empty},
			{engine.Empty, engine.Empty, engine.Empty, engine.Empty},
			{engine.Empty, engine.Empty, engine.Empty, engine.Empty},
		},
		{
			{engine.Empty, engine.Empty, engine.Empty, engine.Empty},
			{engine.Black, engine.Empty, engine.Empty, engine.Empty},
			{engine.Empty, engine.Empty, engine.Empty, engine.Empty},
			{engine.Empty, engine.Empty, engine.Empty, engine.Empty},
		},
		{
			{engine.Empty, engine.Empty, engine.Empty, engine.Empty},
			{engine.Black, engine.Empty, engine.Empty, engine.Empty},
			{engine.Empty, engine.Empty, engine.Empty, engine.Empty},
			{engine.Empty, engine.Empty, engine.Empty, engine.Empty},
		},
		{
			{engine.Empty, engine.Empty, engine.Empty, engine.Empty},
			{engine.Empty, engine.Empty, engine.Empty, engine.Empty},
			{engine.Empty, engine.Empty, engine.Empty, engine.Empty},
			{engine.Empty, engine.Empty, engine.Empty, engine.Empty},
		},
	}
	state.CurrentPlayer = engine.Black
	fmt.Println(ai.SimpleEval(&state, state.CurrentPlayer))
	move := ai.GetNextMove(&state, 3)
	fmt.Println(move)
	_, winner := state.GetWinner()
	fmt.Println(winner.GetName())
}
