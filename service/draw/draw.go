package draw

import (
	"deadlock/config"
	"deadlock/domain/train"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawXJLocks(screen *ebiten.Image) {
	for i := 0; i < config.MAX_TRAINS; i++ {
		drawXJLock(screen, config.JX_LOCKS[i], config.XJ_LOCK_POSITIONS[i][0], config.XJ_LOCK_POSITIONS[i][1])
	}
}

func DrawTrains(screen *ebiten.Image) {
	for i := 0; i < config.MAX_TRAINS; i++ {
		if i%2 == 0 {
			drawXTrain(screen, i, config.TRAINS_DIRECTIONS[i], config.TRAINS_START[i], config.TRAINS_XY_POSITIONS[i])
		} else {
			drawYTrain(screen, i, config.TRAINS_DIRECTIONS[i], config.TRAINS_START[i], config.TRAINS_XY_POSITIONS[i])
		}

	}
}
func DrawTracks(screen *ebiten.Image) {
	for i := 0; i < config.RAILROAD_LENGTH; i++ {
		screen.Set(10+i, 135, config.White)
		screen.Set(185, 10+i, config.White)
		screen.Set(310-i, 185, config.White)
		screen.Set(135, 310-i, config.White)
	}
}

func drawXJLock(screen *ebiten.Image, xjl *train.XJunctionLock, x int, y int) {
	c := config.White
	if xjl.ReservedByTrainId != -1 {
		c = config.Colors[xjl.ReservedByTrainId]
	}
	screen.Set(x-1, y, c)
	screen.Set(x, y-1, c)
	screen.Set(x, y, c)
	screen.Set(x+1, y, c)
	screen.Set(x, y+1, c)
}

func drawXTrain(screen *ebiten.Image, id int, dir config.TrainDirection, start int, yPos int) {
	s := start + (int(dir) * (config.TRAINS[id].LocomotivePosition - config.TRAINS[id].Length))
	e := start + (int(dir) * config.TRAINS[id].LocomotivePosition)

	for i := math.Min(float64(s), float64(e)); i <= math.Max(float64(s), float64(e)); i++ {
		screen.Set(int(i)-int(dir), yPos-1, config.Colors[id])
		screen.Set(int(i), yPos, config.Colors[id])
		screen.Set(int(i)-int(dir), yPos+1, config.Colors[id])
	}
}

func drawYTrain(screen *ebiten.Image, id int, dir config.TrainDirection, start int, xPos int) {
	s := start + (int(dir) * (config.TRAINS[id].LocomotivePosition - config.TRAINS[id].Length))
	e := start + (int(dir) * config.TRAINS[id].LocomotivePosition)
	for i := math.Min(float64(s), float64(e)); i <= math.Max(float64(s), float64(e)); i++ {
		screen.Set(xPos-1, int(i)-int(dir), config.Colors[id])
		screen.Set(xPos, int(i), config.Colors[id])
		screen.Set(xPos+1, int(i)-int(dir), config.Colors[id])
	}
}
