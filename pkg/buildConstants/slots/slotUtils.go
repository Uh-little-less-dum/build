package slots

func StringToSlotId(id string) SlotKey {
	iterator := SlotIdIterator()
	for _, k := range iterator {
		if id == string(k) {
			return k
		}
	}
	return NoSlotApplied
}
