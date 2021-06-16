package slot

// Gets the slot that comes after the current one, nil for end-of-rule
func (s Slot) GetSuccessor() *Slot {
	if s.EoR() {
		return nil
	} else {
		// TODO try slots[s.Label + 1]
		return slots[slotIndex[Index{s.NT,s.Alt,s.Pos+1}]]
	}
}