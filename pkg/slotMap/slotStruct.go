package slot_map

import (
	"sync"

	target_paths "github.com/Uh-little-less-dum/build/pkg/targetPaths"
)

type Slot struct {
	Id       SlotKey
	SubSlots []SubSlot
}

func (s Slot) IsValid() bool {
	return s.Id != NoSlotApplied
}

func (s Slot) WriteOutput(wg *sync.WaitGroup, paths target_paths.TargetPaths) {
	wg.Add(len(s.SubSlots))
	for _, item := range s.SubSlots {
		go func() {
			defer wg.Done()
			item.writeOutput(paths)
		}()
	}
}

func NewSlot(id SlotKey, subSlots []string) Slot {
	var ss = make([]SubSlot, len(subSlots))
	for _, l := range subSlots {
		ss = append(ss, SubSlot{subSlot: l})
	}
	return Slot{Id: id, SubSlots: ss}
}
