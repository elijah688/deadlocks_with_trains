package config

import (
	"deadlock/domain/train"
	"image/color"
)

var (
	TRAINS              = [4]*train.Train{}
	TRAINS_START        = [4]int{10, 10, 310, 310}
	TRAINS_DIRECTIONS   = [4]TrainDirection{EAST, SOUTH, WEST, NORTH}
	TRAINS_XY_POSITIONS = [4]int{135, 185, 185, 135}
	JX_LOCKS            = [4]*train.XJunctionLock{}
	XJ_LOCK_POSITIONS   = [4][2]int{{145, 145}, {175, 145}, {175, 175}, {145, 175}}
	Colors              = [4]color.RGBA{
		{233, 33, 40, 255},
		{78, 151, 210, 255},
		{251, 170, 26, 255},
		{11, 132, 54, 255},
	}

	White = color.RGBA{255, 255, 255, 255}
)

type TrainDirection int

const (
	NORTH                TrainDirection = -1
	EAST                 TrainDirection = 1
	SOUTH                TrainDirection = 1
	WEST                 TrainDirection = -1
	MAX_TRAINS                          = 4
	TRAIN_LENGTH                        = 70
	RAILROAD_LENGTH                     = 300
	MAX_LAYOUT_DIMENSION                = 320
	PIXELS_PER_UNIT                     = 2
)
