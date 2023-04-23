package hierarchy

import (
	"deadlock/domain/train"
	"sort"
	"time"
)

func lockIntersectionsInDistance(id, reserveStart, reserveEnd int, xjs []*train.XJunction) {
	var xjlocks []*train.XJunctionLock
	for _, xj := range xjs {
		if reserveStart <= xj.RailPosition && xj.RailPosition <= reserveEnd && xj.XJunctionLock.ReservedByTrainId != id {
			xjlocks = append(xjlocks, xj.XJunctionLock)
		}
	}

	sort.Slice(xjlocks, func(i, j int) bool {
		return xjlocks[i].Id < xjlocks[j].Id
	})

	for _, xjl := range xjlocks {
		xjl.Mutex.Lock()
		xjl.ReservedByTrainId = id
		time.Sleep(10 * time.Millisecond)
	}
}

func MoveTrain(train *train.Train, distance int, xjs []*train.XJunction) {
	for train.LocomotivePosition < distance {
		train.LocomotivePosition += 1
		for _, xj := range xjs {
			if train.LocomotivePosition == xj.RailPosition {
				lockIntersectionsInDistance(train.Id, xj.RailPosition, xj.RailPosition+train.Length, xjs)
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
