package deadlock

import (
	"deadlock/domain/train"
	"time"
)

func lockIntersectionsInDistance(id, reserveStart, reserveEnd int, xjs []*train.XJunction) {
	var xjlocks []*train.XJunctionLock
	for _, xj := range xjs {
		if reserveStart <= xj.RailPosition && xj.RailPosition <= reserveEnd && xj.XJunctionLock.ReservedByTrainId != id {
			xjlocks = append(xjlocks, xj.XJunctionLock)
		}
	}

	for _, xjl := range xjlocks {
		xjl.ReservedByTrainId = id
		time.Sleep(10 * time.Millisecond)
	}
}

func MoveTrain(train *train.Train, distance int, xjs []*train.XJunction) {
	for train.LocomotivePosition < distance {
		train.LocomotivePosition += 1
		for _, xj := range xjs {
			if train.LocomotivePosition == xj.RailPosition {
				xj.XJunctionLock.Mutex.Lock()
				xj.XJunctionLock.ReservedByTrainId = train.Id
			}
			back := train.LocomotivePosition - train.Length
			if back == xj.RailPosition {
				xj.XJunctionLock.ReservedByTrainId = -1
				xj.XJunctionLock.Mutex.Unlock()
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}
