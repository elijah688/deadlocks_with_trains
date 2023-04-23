package main

import (
	"deadlock/config"
	"deadlock/domain/train"
	"deadlock/service/draw"
	"deadlock/service/hierarchy"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	// Implementation of Update method
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	draw.DrawTracks(screen)
	draw.DrawXJLocks(screen)
	draw.DrawTrains(screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return config.MAX_LAYOUT_DIMENSION, config.MAX_LAYOUT_DIMENSION
}

func main() {
	for i := 0; i < config.MAX_TRAINS; i++ {
		config.TRAINS[i] = train.NewTrain(i, config.TRAIN_LENGTH)
		config.JX_LOCKS[i] = train.NewXJunctionLock(i)
	}
	for i := 0; i < config.MAX_TRAINS; i++ {
		go hierarchy.MoveTrain(config.TRAINS[i], 300, []*train.XJunction{train.NewXJunction(125, config.JX_LOCKS[i]), train.NewXJunction(175, config.JX_LOCKS[(i+1)%4])})
	}

	ebiten.SetWindowSize(320*3, 320*3)
	ebiten.SetWindowTitle("XJLocks ### CHOOO... CHOOO!")
	if err := ebiten.RunGame(new(Game)); err != nil {
		log.Fatal(err)
	}
}
