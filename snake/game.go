package snake

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 600
	ScreenHeight = 600
	boardRaws    = 20
	boardCols    = 20
)

var (
	backgroundColor = color.RGBA{R: 50, G: 100, B: 50, A: 50}
	snakeColor      = color.RGBA{R: 200, G: 50, B: 150, A: 150}
	foodColor       = color.RGBA{R: 200, G: 200, B: 50, A: 150}
)

type Game struct {
	input *Input
	board *Board
}

func NewGame() *Game {
	return &Game{
		input: NewInput(),
		board: NewBorad(boardRaws, boardCols),
	}
}

func (g *Game) Layout(outsideWith, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
func (g *Game) Update() error {
	return g.board.Update(g.input)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	if g.board.gameOver {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Game Over. Score: %d", g.board.points))
	} else {
		width := ScreenHeight / boardRaws

		for _, p := range g.board.snake.body {
			ebitenutil.DrawRect(screen, float64(p.y*width), float64(p.x*width), float64(width), float64(width),
				snakeColor)
		}
		if g.board.food != nil {
			ebitenutil.DrawRect(screen, float64(g.board.food.y*width), float64(g.board.food.x*width), float64(width),
				float64(width), foodColor)
		}
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.board.points))
	}
}
