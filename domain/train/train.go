package train

import (
	"sync"
)

type Train struct {
	Id                 int
	Length             int
	LocomotivePosition int
}

func NewTrain(Id, Length int) *Train {
	t := new(Train)
	t.Id = Id
	t.Length = Length

	return t
}

type XJunctionLock struct {
	Id                int
	Mutex             *sync.Mutex
	ReservedByTrainId int
}

func NewXJunctionLock(Id int) *XJunctionLock {
	xjl := new(XJunctionLock)
	xjl.Id = Id
	xjl.ReservedByTrainId = -1
	xjl.Mutex = new(sync.Mutex)

	return xjl
}

type XJunction struct {
	RailPosition  int
	XJunctionLock *XJunctionLock
}

func NewXJunction(RailPosition int, XJunctionLock *XJunctionLock) *XJunction {
	xj := new(XJunction)
	xj.RailPosition = RailPosition
	xj.XJunctionLock = XJunctionLock
	return xj
}
