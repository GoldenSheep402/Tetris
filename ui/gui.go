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
		"Just write for fun." +
		"Use e to turn the tetromino \n" +
		"(You should decide the direction before move left or right)\n" +
		"Press a to move left.\n" +
		"Press d to move right.\n" +
		"Enter to drop.\n",
	)
}

func PrintBoard(g *define.Game) {
	// 打印游戏板
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	PrintInfo()
	fmt.Printf("|------------------------------|[SCORE:%d | LEVAL:%d]\n", g.Score, g.Level)
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
