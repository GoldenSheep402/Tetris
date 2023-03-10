package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
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
	dropHeight := 999

	// 下落高度
	for x := 0; x < 15; x++ {
		top := 0
		button := 40
		for y := 0; y < 3; y++ {
			if g.board[y][x] == 1 {
				if top < y {
					top = y
				}
			}
		}
		for y := 39; y >= 4; y-- {
			if g.board[y][x] == 1 {
				if button > y {
					button = y
				}
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
				g.board[y+dropHeight-1][x] = 1
				g.board[y][x] = 0
			}
		}
	}
}

func (g *Game) ClearFullRows() {
	// 消除满行
	clearRows := 0
	flag := true
	
	for y := 39; y >= 0; y-- {
		for x := 0; x < 15; x++ {
			if g.board[y][x] == 1 {
				flag = false
				break
			}
			if flag {
				clearRows++
				g.score += 100
			}
		}
	}

	for i := 0; i < clearRows; i++ {
		for x := 0; x < 15; x++ {
			for y := 39; y >= 0; y-- {
				if g.board[y][x] == 1 && (y+i+1) < 40 {
					g.board[y+i+1][x] = 1
					g.board[y][x] = 0
				}
			}
		}
	}
}

func (g *Game) Start() {
	for {
		g.NewTetrominoIn()
		g.PrintBoard()

		// 等待用户按下回车键
		reader := bufio.NewReader(os.Stdin)
		_, _, err := reader.ReadRune()

		if err != nil {

			break
		} else {
			// 检查操作系统是否是 Windows
			if runtime.GOOS == "windows" {
				// 在 Windows 上清除控制台屏幕的内容
				cmd := exec.Command("cmd", "/c", "cls")
				cmd.Stdout = os.Stdout
				cmd.Run()

				// 执行 DropTetromino 函数
				g.DropTetromino()
				g.ClearFullRows()
			}
		}
	}
	fmt.Println("Game Over!")
}

func (g *Game) PrintBoard() {
	// 打印游戏板
	fmt.Print("--------------------------------\n")
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
	game := Game{
		board:  make([][]int, 40),
		active: make([][]int, 4),
		score:  0,
		level:  1,
	}
	game.BoardInit()
	game.Start()
}
