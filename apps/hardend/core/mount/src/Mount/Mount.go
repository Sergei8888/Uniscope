package mount

import (
	"core/mount/src/Coordinates/Earth"
	"time"
)

type Mount interface {
	GetSkyCoordinates(skyCoordinatesType int, precise bool) (vertRotation, horizRotation *Rotation, err error)
	GotoSkyCoordinates(skyCoordinatesType int, vertRotation, horizRotation *Rotation) error
	GetEarthCoordinates() (*Earth.Location, error)
	SetEarthCoordinates(loc *Earth.Location) error
	SlewBySpeed(speed int) error
	SlewByRate(rate int) error
	SyncBySkyObject(skyCoordinatesType int, precise bool, vertRotation, horizRotation *Rotation) error
	GetTime() (*time.Time, error)
	SetTime(t *time.Time) error
	SetTrackingMode(mode int) error
	GetTrackingMode() (mode int, err error)
	GetInfo() *Info
}
