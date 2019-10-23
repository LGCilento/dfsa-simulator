package dfsa

import (
	"frame"
	"math"
	"time"
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
	EstimationTime  float64
	EstimationError float64 //adicionado
}

func (s *Simulator) Run() (result SimulationResult) {
	var qntIterations = 0
	var estimationTime = 0
	currentFrame := frame.Frame{
		Size: s.InitialFrameSize,
	}
	backlog := s.InitialTagsLen
	for {
		currentFrame.TransmitTags(s.InitialTagsLen)
		qntIterations++
		var estimationTimeInit = time.Now()
		result.computeFrame(currentFrame)
		var estimationTimeFinal = time.Since(estimationTimeInit)
		estimationTime += int(estimationTimeFinal)
		backlog -= currentFrame.SuccessfulSlots
		if backlog <= 0 {
			result.EstimationTime = float64(estimationTime / qntIterations)
			result.EstimationError = result.EstimationError / float64(qntIterations)
			break
		}
		currentFrame = s.Estimator.GetNextFrame(currentFrame)
		//var estimationTimeFinal = time.Since(estimationTimeInit)
	}
	return
}

func (r *SimulationResult) computeFrame(fr frame.Frame) {
	r.SlotsSum += fr.Size
	r.SuccessfulSlots += fr.SuccessfulSlots
	r.CollisionSlots += fr.CollisionSlots
	r.EmptySlots += fr.EmptySlots
	r.EstimationError += math.Abs(float64(fr.Backlog - fr.Size)) //adicionado
}
