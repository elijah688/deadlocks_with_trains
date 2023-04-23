package arbiter

import (
	"deadlock/domain/train"
	"sync"
	"time"
)

var (
	mutex = new(sync.Mutex)
	cond  = sync.NewCond(mutex)
)

func allFree(xjls []*train.XJunctionLock) bool {

	for _, xjl := range xjls {
		if xjl.ReservedByTrainId != -1 {
			return false
		}
	}

	return true

}
func lockIntersectionsInDistance(id, reserveStart, reserveEnd int, xjs []*train.XJunction) {
	var xjlocks []*train.XJunctionLock
	for _, xj := range xjs {
		if reserveStart <= xj.RailPosition && xj.RailPosition <= reserveEnd && xj.XJunctionLock.ReservedByTrainId != id {
			xjlocks = append(xjlocks, xj.XJunctionLock)
		}
	}

	cond.L.Lock()
	for !allFree(xjlocks) {
		cond.Wait()
	}

	for _, xjl := range xjlocks {
		xjl.ReservedByTrainId = id
		time.Sleep(10 * time.Millisecond)
	}

	cond.L.Unlock()
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
				cond.L.Lock()
				xj.XJunctionLock.ReservedByTrainId = -1
				cond.Broadcast()
				cond.L.Unlock()
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}
