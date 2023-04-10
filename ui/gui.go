package ui

import (
	"fmt"
	"os"
	"os/exec"
	"tetris/define"
)

func PrintInfo() {
	fmt.Print("Tetris Game(SB version)\n" +
		"Author: GS\n" + "" +
		"Just write for fun.\n" +
		"[Select your choice then ENTER]\n" +
		"Use e to turn the tetromino \n" +
		"Press a to move left.\n" +
		"Press d to move right.\n" +
		"Press s to drop.\n",
	)
}

func PrintBoard(g *define.Game) {
	// 打印游戏板
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	PrintInfo()
	fmt.Printf("\n ┌───────────────────────────┐\n │SCORE:%-7d│LEVEL:%-7d│\n └───────────────────────────┘\n|------------------------------|\n", g.Score, g.Level)
	for i := range g.Board {
		fmt.Print("|")
		if i == 4 {
			fmt.Print("------------------------------|\n|")
		}
		for j := range g.Board[i] {
			if g.Board[i][j] == 0 {
				fmt.Print("  ")
			} else {
				fmt.Print("■ ")
			}
		}
		fmt.Printf("|\n")
	}
	fmt.Print("|------------------------------|\n")
}
