package logic

import (
	"math/rand"
	"tetris/define"
	"tetris/ui"
	"time"
)

const (
	fps            = 30
	blockMoveDelay = time.Millisecond * 200
)

type Game struct {
	board        [][]int
	currBlock    [][]int
	nextBlock    [][]int
	currBlockX   int
	currBlockY   int
	currBlockRot int
}

func (g *Game) Init() {
	g.board = make([][]int, define.BoardHeight)
	for i := 0; i < define.BoardHeight; i++ {
		g.board[i] = make([]int, define.BoardWidth)
	}
}

func (g *Game) Start() {
	g.Init()

	// 在顶部随机放置一个方块
	g.nextBlock, _ = GetRandomBlock()
	g.currBlock = g.nextBlock
	g.currBlockX = define.BoardWidth/2 - define.BlockSize*2
	g.currBlockY = define.TopPadding
	g.nextBlock, _ = GetRandomBlock()

	for {
		ui.DrawBoard()

		// 在顶部随机放置下一个方块
		if g.currBlock == nil {
			g.currBlock = g.nextBlock
			g.currBlockX = define.BoardWidth/2 - define.BlockSize*2
			g.currBlockY = define.TopPadding
			g.nextBlock, _ = GetRandomBlock()
		}

		time.Sleep(time.Second / fps)
	}
}

func (g *Game) Drop() {
	// 将当前方块放到底部
}

func (g *Game) MoveLeft() {
	// 将当前方块向左移动
}

func (g *Game) MoveRight() {
	// 将当前方块向右移动
}

func (g *Game) Rotate() {
	// 将当前方块旋转
}

func GameStart() {
	game := Game{}
	game.Start()
}

func GetRandomBlock() ([][]int, int) {
	index := rand.Intn(len(define.Tetrominoes))
	return define.Tetrominoes[index], index
}
