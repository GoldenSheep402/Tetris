package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"tetris/define"
	"time"
)

type Game struct {
	board  [][]int
	active [][]int
	score  int
	level  int
}

func (g *Game) BoardInit() {
	// 初始化游戏板
	g.board = make([][]int, 40)
	for i := range g.board {
		g.board[i] = make([]int, 15)
	}
}

func (g *Game) NewTetrominoIn() {
	// 生成新的俄罗斯方块
	nextTetromino := RandomTetromino()
	for i := range nextTetromino {
		for j := range nextTetromino[i] {
			if nextTetromino[i][j] == 1 {
				g.board[i][j] = 1
			}
		}
	}
}

func (g *Game) DropTetromino() {
	// 俄罗斯方块下落
	top := 0
	button := 39
	dropHeight := 999

	// 下落高度
	for x := 0; x < 15; x++ {
		for y := 0; y < 4; y++ {
			if g.board[y][x] == 1 {
				top = y
			}
		}

		for y := 39; y >= 36; y-- {
			if g.board[y][x] == 1 {
				button = y
			}
		}

		if dropHeight > (button - top) {
			dropHeight = button - top
		}
	}

	// 下落
	for x := 0; x < 15; x++ {
		for y := 3; y >= 0; y-- {
			if g.board[y][x] == 1 {
				g.board[y+dropHeight][x] = 1
				g.board[y][x] = 0
			}
		}
	}
	g.PrintBoard()
}

func (g *Game) Start() {
	for {
		g.NewTetrominoIn()

		g.PrintBoard()
		// 等待用户按下回车键
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			// 发生错误，退出循环
			break
		}
		if char == '\n' {
			// 用户按下回车键，退出循环
			g.DropTetromino()
		}
		//g.PrintBoard()
	}
}

func (g *Game) PrintBoard() {
	// 打印游戏板

	for i := range g.board {
		fmt.Print("|")
		for j := range g.board[i] {
			if g.board[i][j] == 0 {
				fmt.Print("  ")
			} else {
				fmt.Print("■ ")
			}
		}
		fmt.Print("|\n")
	}
}

func RandomTetromino() [][]int {
	// 对随机数生成器进行种子随机化
	rand.Seed(time.Now().UnixNano())

	// 从Tetrominoes数组中随机选择一个形状
	shape := define.Tetrominoes[rand.Intn(len(define.Tetrominoes))]

	// 复制所选形状的矩阵
	tetromino := make([][]int, len(shape))
	for i := range shape {
		tetromino[i] = make([]int, len(shape[i]))
		copy(tetromino[i], shape[i])
	}

	return tetromino
}

func main() {
	game := Game{}
	game.BoardInit()
	game.Start()
}
