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
