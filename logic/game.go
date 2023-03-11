package logic

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"tetris/define"
	"tetris/ui"
	"time"
)

func BoardInit(g *define.Game) {
	// 初始化游戏板
	g.Board = make([][]int, define.HEIGHT)
	for i := range g.Board {
		g.Board[i] = make([]int, define.WIDTH)
	}
}

func Start(g *define.Game) {
	for {
		NewTetrominoIn(g)
		ui.PrintBoard(g)
		CheckGameOver(g)

		var x int = 0

		for {
			reader := bufio.NewReader(os.Stdin)
			input, _, err := reader.ReadRune()
			if err != nil {
				break
			} else if input == '\r' { // 回车键
				break
			} else {
				Move(g, string(input), &x)
				ui.PrintBoard(g)
			}
		}

		DropTetromino(g)
		ClearFullRows(g)
	}
}

func NewTetrominoIn(g *define.Game) {
	// 生成新的俄罗斯方块
	nextTetromino := RandomTetromino()
	for i := range nextTetromino {
		for j := range nextTetromino[i] {
			if nextTetromino[i][j] == 1 {
				g.Board[i][j] = 1
			}
		}
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

func CheckGameOver(g *define.Game) {
	// 检查游戏是否结束
	for x := 0; x < define.WIDTH; x++ {
		if g.Board[4][x] == 1 {
			fmt.Printf("[Game Over!]\n[Your Score: %d]\nThank you for playing!", g.Score)
			os.Exit(0)
		}
	}
}

func Move(g *define.Game, direction string, driectX *int) {
	switch direction {
	case "a":
		for y := 0; y < 4; y++ {
			if g.Board[y][0] == 1 {
				return
			}
		}
		*driectX--
		// 向左移动
		for y := 0; y <= 3; y++ {
			for x := 1; x < define.WIDTH; x++ {
				if g.Board[y][x] == 1 {
					if g.Board[y][x-1] == 0 {
						g.Board[y][x-1] = 1
						g.Board[y][x] = 0
					}
				}
			}
		}
	case "d":
		for y := 0; y < 4; y++ {
			if g.Board[y][define.WIDTH-1] == 1 {
				return
			}
		}
		*driectX++
		// 向右移动
		for y := 0; y <= 3; y++ {
			for x := define.WIDTH - 2; x >= 0; x-- {
				if g.Board[y][x] == 1 {
					if g.Board[y][x+1] == 0 {
						g.Board[y][x+1] = 1
						g.Board[y][x] = 0
					}
				}
			}
		}
	case "e":
		// 旋转
		rotateClockwise(g.Board, 0, *driectX)
	}
}

func rotateClockwise(matrix [][]int, row int, col int) {
	// 防止溢出
	if col+4 > define.WIDTH || row+4 > define.HEIGHT {
		return
	}

	// 定义临时矩阵
	temp := make([][]int, 4)
	for i := range temp {
		temp[i] = make([]int, 4)
	}

	// 拷贝要旋转的4x4矩阵到临时矩阵中
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			temp[i][j] = matrix[row+i][col+j]
		}
	}

	// 对临时矩阵进行旋转操作
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			matrix[row+i][col+j] = temp[j][3-i]
		}
	}
}

func DropTetromino(g *define.Game) {
	// 俄罗斯方块下落
	dropHeight := 999

	// 下落高度
	for x := 0; x < define.WIDTH; x++ {
		top := 0
		button := define.HEIGHT
		for y := 0; y <= 3; y++ {
			if g.Board[y][x] == 1 {
				if top < y {
					top = y
				}
			}
			for y := define.HEIGHT - 1; y > 4; y-- {
				if g.Board[y][x] == 1 {
					if button > y {
						button = y
					}
				}
				if (top != 0) && (dropHeight > (button - top)) {
					dropHeight = button - top
				}
			}
		}
	}

	// 下落
	for x := 0; x < define.WIDTH; x++ {
		for y := 3; y >= 0; y-- {
			if g.Board[y][x] == 1 {
				g.Board[y+dropHeight-1][x] = 1
				g.Board[y][x] = 0
			}
		}
	}
}

func ClearFullRows(g *define.Game) {
	clearRows := 0
	flag := true

	for y := define.HEIGHT - 1; y >= 0; y-- {
		flag = true
		for x := 0; x < define.WIDTH; x++ {
			if g.Board[y][x] == 0 {
				flag = false
				break
			}
		}
		if flag {

			// 去除被消去的行
			for x := 0; x < define.WIDTH; x++ {
				g.Board[y][x] = 0
			}

			Drop(g, y)

			g.Level++
			clearRows++
			g.Score += 100
		}
	}
}

func Drop(g *define.Game, startLine int) {
	for x := 0; x < define.WIDTH; x++ {
		var dropHeight int = 0
		top := startLine
		for y := startLine; y >= 0; y-- {
			if g.Board[y][x] == 1 {
				top = y
				break
			}
		}

		for y := top; y < define.HEIGHT; y++ {
			if y == define.HEIGHT-1 || g.Board[y+1][x] == 1 {
				dropHeight = y - top
				break
			}
		}

		for y := top; y > 5; y-- {
			g.Board[y+dropHeight][x] = g.Board[y][x]
			g.Board[y][x] = 0
		}
	}
}
