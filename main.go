package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
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
	g.board = make([][]int, define.HEIGHT)
	for i := range g.board {
		g.board[i] = make([]int, define.WIDTH)
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
	for x := 0; x < define.WIDTH; x++ {
		top := 0
		button := define.HEIGHT
		for y := 0; y <= 3; y++ {
			if g.board[y][x] == 1 {
				if top < y {
					top = y
				}
			}
		}
		for y := define.HEIGHT - 1; y >= 4; y-- {
			if g.board[y][x] == 1 {
				if button > y {
					button = y
				}
			}
		}
		if (top != 0) && (dropHeight > (button - top)) {
			dropHeight = button - top
		}
	}

	// 下落
	for x := 0; x < define.WIDTH; x++ {
		for y := 3; y >= 0; y-- {
			if g.board[y][x] == 1 {
				g.board[y+dropHeight-1][x] = 1
				g.board[y][x] = 0
			}
		}
	}
}

func (g *Game) ClearFullRows() {
	clearRows := 0
	flag := true

	for y := define.HEIGHT - 1; y >= 0; y-- {
		flag = true
		for x := 0; x < define.WIDTH; x++ {
			if g.board[y][x] == 0 {
				flag = false
				break
			}
		}
		if flag {

			// 去除被消去的行
			for x := 0; x < define.WIDTH; x++ {
				g.board[y][x] = 0
			}

			g.Drop(y)

			g.level++
			clearRows++
			g.score += 100
		}
	}
}

// 传入被消去的行标
func (g *Game) Drop(startLine int) {
	for x := 0; x < define.WIDTH; x++ {
		var dropHeight int = 0
		top := startLine
		for y := startLine; y >= 0; y-- {
			if g.board[y][x] == 1 {
				top = y
				break
			}
		}

		for y := top; y < define.HEIGHT; y++ {
			if y == define.HEIGHT-1 || g.board[y+1][x] == 1 {
				dropHeight = y - top
				break
			}
		}

		for y := top; y > 5; y-- {
			g.board[y+dropHeight][x] = g.board[y][x]
			g.board[y][x] = 0
		}
	}
}

func (g *Game) Move(direction string) {
	location := 0
	switch direction {
	case "a":
		location--
		// 向左移动
		for y := 0; y <= 3; y++ {
			for x := 1; x < define.WIDTH; x++ {
				if g.board[y][x] == 1 {
					if g.board[y][x-1] == 0 {
						g.board[y][x-1] = 1
						g.board[y][x] = 0
					}
				}
			}
		}
	case "d":
		location++
		// 向右移动
		for y := 0; y <= 3; y++ {
			for x := define.WIDTH - 2; x >= 0; x-- {
				if g.board[y][x] == 1 {
					if g.board[y][x+1] == 0 {
						g.board[y][x+1] = 1
						g.board[y][x] = 0
					}
				}
			}
		}
	case "e":
		// 旋转
		rotateClockwise(g.board, 0, location)
	}
}

func rotateClockwise(matrix [][]int, row, col int) {
	// 矩阵转置
	for i := row; i < row+4; i++ {
		for j := col; j < col+4; j++ {
			if i < j {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
	}

	// 每行反转
	for i := row; i < row+4; i++ {
		for j := col; j < col+2; j++ {
			matrix[i][j], matrix[i][col+3-j] = matrix[i][col+3-j], matrix[i][j]
		}
	}
}

func (g *Game) CheckGameOver() {
	// 检查游戏是否结束
	for x := 0; x < define.WIDTH; x++ {
		if g.board[4][x] == 1 {
			fmt.Printf("Game Over!\n[Your Score: %d]", g.score)
			os.Exit(0)
		}
	}
}

func (g *Game) Start() {
	for {
		g.NewTetrominoIn()
		g.PrintBoard()
		g.CheckGameOver()

		for {
			reader := bufio.NewReader(os.Stdin)
			input, _, err := reader.ReadRune()
			if err != nil {
				break
			} else if input == '\r' { // 回车键
				break
			} else {
				g.Move(string(input))
				g.PrintBoard()
			}
		}

		g.DropTetromino()
		g.ClearFullRows()
	}
}

func (g *Game) PrintBoard() {
	// 打印游戏板
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Printf("|------------------------------|[SCORE:%d | LEVAL:%d]\n", g.score, g.level)
	for i := range g.board {
		fmt.Print("|")
		if i == 4 {
			fmt.Print("------------------------------|\n|")
		}
		for j := range g.board[i] {
			if g.board[i][j] == 0 {
				fmt.Print("  ")
			} else {
				fmt.Print("■ ")
			}
		}
		fmt.Printf("|\n")
	}
	fmt.Print("|------------------------------|\n")
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
		level: 1,
	}
	game.BoardInit()
	game.Start()
}
