package dfsa

import (
	"frame"
	"math"
)

type Simulator struct {
	Estimator        Estimator
	InitialTagsLen   int
	InitialFrameSize int
}

type SimulationResult struct {
	SlotsSum        int
	EmptySlots      int
	SuccessfulSlots int
	CollisionSlots  int
	EstimationError float64 //adicionado
}

func (s *Simulator) Run() (result SimulationResult) {
	currentFrame := frame.Frame{
		Size: s.InitialFrameSize,
	}
	backlog := s.InitialTagsLen
	for {
		currentFrame.TransmitTags(s.InitialTagsLen)
		result.computeFrame(currentFrame)
		backlog -= currentFrame.SuccessfulSlots
		if backlog <= 0 {
			break
		}
		currentFrame = s.Estimator.GetNextFrame(currentFrame)
	}
	return
}

func (r *SimulationResult) computeFrame(fr frame.Frame) {
	r.SlotsSum += fr.Size
	r.SuccessfulSlots += fr.SuccessfulSlots
	r.CollisionSlots += fr.CollisionSlots
	r.EmptySlots += fr.EmptySlots
	r.EstimationError += math.Abs(float64(fr.CompetingTags + fr.SuccessfulSlots - fr.Size)) //adicionado
}
