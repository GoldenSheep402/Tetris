package main

import (
	"tetris/define"
	"tetris/logic"
)

func main() {
	game := define.Game{
		Level: 1,
	}
	logic.BoardInit(&game)
	logic.Start(&game)
}
