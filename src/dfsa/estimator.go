package dfsa

import (
	"frame"
	"math"
)

type Estimator interface {
	GetNextFrame(frame.Frame) frame.Frame
}

type LowerBound struct{}

func (lb *LowerBound) GetNextFrame(currentFrame frame.Frame) (nextFrame frame.Frame) {
	nextFrame.Size = 2 * currentFrame.CollisionSlots
	return
}

type Schoute struct{}

func (lb *Schoute) GetNextFrame(currentFrame frame.Frame) (nextFrame frame.Frame) {
	nextFrame.Size = int(math.Ceil(2.39 * float64(currentFrame.CollisionSlots)))
	return
}

type EomLee struct{}

func (e *EomLee) GetNextFrame(currentFrame frame.Frame) (nextFrame frame.Frame) {
	var beta, prevY, num, den, frac float64
	currY := 2.0
	epsilon := 1e-3
	for {
		prevY = currY
		beta = float64(currentFrame.Size) / ((prevY * float64(currentFrame.CollisionSlots)) + float64(currentFrame.SuccessfulSlots))
		frac = math.Pow(math.E, -(1 / beta))
		num = 1 - frac
		den = beta * (1 - (1+(1/beta))*frac)
		currY = num / den
		if math.Abs(currY-prevY) < epsilon {
			break
		}
	}
	nextFrame.Size = int(math.Ceil(currY * float64(currentFrame.CollisionSlots)))
	return
}

type Vogt struct{}

func (e *Vogt) GetNextFrame(currentFrame frame.Frame) (nextFrame frame.Frame) {
	e0, e1 := 0.0, -1.0
	i := float64(currentFrame.SuccessfulSlots + (2 * currentFrame.CollisionSlots))
	a0, a1, a2 := 0.0, 0.0, 0.0
	n := i
	for {
		t := 1 - float64(1/currentFrame.Size)
		a0 = math.Pow(t, n)
		a1 = (n * a0) / (float64(currentFrame.Size) * t)
		a2 = 1 - (a1 + a0)
		a0 = (a0 * float64(currentFrame.Size)) - float64(currentFrame.EmptySlots)
		a1 = (a1 * float64(currentFrame.Size)) - float64(currentFrame.SuccessfulSlots)
		a2 = (a2 * float64(currentFrame.Size)) - float64(currentFrame.CollisionSlots)
		a0, a1, a2 = a0*a0, a1*a1, a2*a2
		e0 = e1
		e1 = math.Sqrt(a0 + a1 + a2)
		if n == i {
			e0 = e1 + 1
		}
		n = n + 1
		if e1 < e0 {
			break
		}
	}

	nextFrame.Size = int(n) - 1 - currentFrame.Size
	return
}
