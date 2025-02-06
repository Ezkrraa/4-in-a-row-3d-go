package ai_test

import (
	"cubes/main/ai"
	"cubes/main/engine"
	"math"
	"testing"
)

func TestWinMove(t *testing.T) {
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
			{engine.Empty, engine.Empty, engine.Empty, engine.Empty},
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
	ai_move := ai.GetNextMove(&state, 1)
	if ai_move.X != 3 || ai_move.Y != 0 {
		t.Fatalf("loss not detected properly, move taken: {%d, %d}", ai_move.X, ai_move.Y)
	}
}

func TestEvaluateLoss(t *testing.T) {
	state := engine.GameState{
		Board: [4][4][4]engine.FieldState{
			{
				{engine.White, engine.White, engine.White, engine.White},
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
				{engine.Black, engine.Empty, engine.Empty, engine.Empty},
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
		},
		CurrentPlayer: engine.Black,
	}
	eval := ai.EvaluateState(&state, 2, -math.MaxFloat32, math.MaxFloat32, state.CurrentPlayer, state.CurrentPlayer)
	if eval > -1_000 {
		t.Fatalf("Eval: %f", eval)
	}
}

func TestMovingAndWinning(t *testing.T) {
	state := engine.GameState{
		Board: [4][4][4]engine.FieldState{
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
		},
		CurrentPlayer: engine.White,
	}
	state = *state.GetMovedClone(engine.MoveCoordinate{X: 3, Y: 0})
	finished, winner := state.GetWinner()
	if !finished {
		t.Fatal("Losses not detected properly")
	} else if winner != state.CurrentPlayer.Flip() {
		t.Fatalf("Winner was %s", winner.GetName())
	} else {
		t.Logf("Winner was %s", winner.GetName())
	}
}
