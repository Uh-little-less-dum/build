package slots

type Slot struct {
	Id       SlotKey
	SubSlots []string
}

func (s Slot) IsValid() bool {
	return s.Id != NoSlotApplied
}

func NewSlot(id SlotKey, subSlots []string) Slot {
	return Slot{Id: id, SubSlots: subSlots}
}
