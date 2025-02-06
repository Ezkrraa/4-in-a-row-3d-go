package ui

import (
	"cubes/main/engine"
	"errors"
	"fmt"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

var ResetColor = "\033[0m"
var WhiteBgColor = "\033[47m"
var BlackColor = "\033[30m"

func selectOption(title string, options []string) (int, error) {
	selected := 0
	for {
		if selected < 0 {
			selected = 0
		} else if selected > len(options)-1 {
			selected = len(options) - 1
		}
		clearTerminal()
		fmt.Println(title)
		fmt.Println(selected)
		for i := range options {
			if selected == i {
				fmt.Print(" > " + options[i] + "\r\n")
			} else {
				fmt.Print(" - " + options[i] + "\r\n")
			}
		}
		err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {
			switch key.Code {
			case keys.Up:
				selected--
			case keys.Down:
				selected++
			case keys.Enter:
				fallthrough
			case keys.Space:
				return true, errors.New("select")
			case keys.CtrlC:
				fallthrough
			case keys.Esc:
				return true, errors.New("quit")
			default:
				break
			}
			return true, nil // Return false to continue listening
		})
		if err != nil {
			if err.Error() == "select" {
				return selected, nil
			} else if err.Error() == "quit" {
				return selected, err
			}
			panic("Unknown error: " + err.Error())
		}
	}
}

func selectMove(state *engine.GameState) (engine.MoveCoordinate, error) {
	selected := engine.BoardCoordinate{X: 2, Y: 0, Z: 0}
	for {
		clearTerminal()
		selected.X = min(3, max(0, selected.X))
		selected.Y = min(3, max(0, selected.Y))
		selected.Z = min(3, max(0, selected.Z))

		// fmt.Printf("X: %d, Y: %d, Z: %d\n", int(selected.X), int(selected.Y), int(selected.Z))
		for z := int8(0); z < 4; z++ {
			for y := int8(0); y < 4; y++ {
				for x := int8(0); x < 4; x++ {
					if selected.X == x && selected.Y == y && selected.Z == z {
						fmt.Print(WhiteBgColor + BlackColor)
						fmt.Print(state.Board[z][y][x].GetString())
						fmt.Print(ResetColor)
					} else {
						fmt.Print(state.Board[z][y][x].GetString())
					}
					fmt.Print(" ")
				}
				fmt.Print("\r\n")
			}
			fmt.Print("\r\n")
		}

		err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {
			switch key.Code {
			case keys.Up:
				if selected.Y < 1 {
					if selected.Z > 0 {
						selected.Z--
						selected.Y = 3
					}
				} else {
					selected.Y--
				}
			case keys.Down:
				if selected.Y > 2 && selected.Z < 3 {
					selected.Z++
					selected.Y = 0
				} else {
					selected.Y++
				}
			case keys.Left:
				selected.X--
			case keys.Right:
				selected.X++
			case keys.Enter:
				fallthrough
			case keys.Space:
				return true, errors.New("selected")
			case keys.CtrlC:
				fallthrough
			case keys.Esc:
				return true, errors.New("quit")
			default:
				break
			}
			return true, nil // Return false to continue listening
		})
		if err != nil {
			if err.Error() == "selected" {
				return engine.MoveCoordinate{X: selected.X, Y: selected.Y}, nil
			} else if err.Error() == "quit" {
				return engine.MoveCoordinate{X: selected.X, Y: selected.Y}, err
			}
		}
	}
}

func printState(state *engine.GameState) {
	for z := int8(0); z < 4; z++ {
		for y := int8(0); y < 4; y++ {
			for x := int8(0); x < 4; x++ {
				fmt.Print(state.Board[z][y][x].GetString())
				fmt.Print(" ")
			}
			fmt.Print("\r\n")
		}
		fmt.Print("\r\n")
	}
}

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}
