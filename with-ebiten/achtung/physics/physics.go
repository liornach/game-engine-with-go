package physics

import (
	"math"
	"time"
)

func EstimateFutureLocation(sv StateVector, t time.Duration) StateVector {
	l := sv.location
	v := sv.velocity

	l.X += v.X * t.Seconds()
	l.Y += v.Y * t.Seconds()

	return NewStateVector(l, v)
}

func Rotate(sv StateVector, rad float64) StateVector {
	cos := math.Cos(rad)
	sin := math.Sin(rad)

	v := sv.velocity
	oldx, oldy := v.X, v.Y

	v.X = oldx*cos - oldy*sin
	v.Y = oldx*sin + oldy*cos

	return NewStateVector(sv.location, v)
}
