package frame

import "math/rand"

type Frame struct {
	Size            int
	SuccessfulSlots int
	CollisionSlots  int
	EmptySlots      int
	CompetingTags   int
}

func (f *Frame) TransmitTags(tagsLen int) {
	transmissionOrder := f.buildTransmissionOrder(tagsLen)
	slots := make([]int, f.Size)
	for _, o := range transmissionOrder {
		slots[o] += 1
	}

	for _, tagsTransmitted := range slots {
		if tagsTransmitted == 0 {
			f.EmptySlots += 1
		} else if tagsTransmitted == 1 {
			f.SuccessfulSlots += 1
		} else {
			f.CollisionSlots += 1
			f.CompetingTags += tagsTransmitted
		}
	}
}

// buildTransmissionOrder gives each tag a slot to transmit it
func (f *Frame) buildTransmissionOrder(tagsLen int) []int {
	order := make([]int, tagsLen)
	for i := 0; i < tagsLen; i++ {
		allocatedSlot := rand.Intn(f.Size)
		order[i] = allocatedSlot
	}
	return order
}
