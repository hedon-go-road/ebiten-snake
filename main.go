package main

import (
	"log"

	"ebiten-snake/snake"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := snake.NewGame()

	ebiten.SetWindowSize(snake.ScreenWidth, snake.ScreenHeight)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
