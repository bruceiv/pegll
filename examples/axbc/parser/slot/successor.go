package slot

import(
	"axbc/token"
)

// Gets the slot that comes after the current one, nil for end-of-rule
func (s Slot) GetSuccessor() *Slot {
	if s.EoR() {
		return nil
	} else {
		// TODO try slots[s.Label + 1]
		return slots[slotIndex[Index{s.NT,s.Alt,s.Pos+1}]]
	}
}

// Checks nullability of a slot
func (l Label) IsNullable() bool {
	return nullable[l]
}

// Checks FIRST set of a slot
func (l Label) FirstContains(typ token.Type) bool {
	return firstT[l][typ]
}

// map of slot labels to their nullability
var nullable = []bool {
	// AorB : ∙As
	true,
	// AorB : As ∙
	true,
	// AorB : ∙a b
	false,
	// AorB : a ∙b
	false,
	// AorB : a b ∙
	true,
	// As : ∙a As
	false,
	// As : a ∙As
	true,
	// As : a As ∙
	true,
	// As : ∙
	true,
	// AxBC : ∙AorB c
	false,
	// AxBC : AorB ∙c
	false,
	// AxBC : AorB c ∙
	true,
}

// map of slot labels to tokens to containment in the FIRST set
var firstT = []map[token.Type]bool {
	// AorB : ∙As
	{
		token.T_0: true,
	},
	// AorB : As ∙
	{
	},
	// AorB : ∙a b
	{
		token.T_0: true,
	},
	// AorB : a ∙b
	{
		token.T_1: true,
	},
	// AorB : a b ∙
	{
	},
	// As : ∙a As
	{
		token.T_0: true,
	},
	// As : a ∙As
	{
		token.T_0: true,
	},
	// As : a As ∙
	{
	},
	// As : ∙
	{
	},
	// AxBC : ∙AorB c
	{
		token.T_0: true,
		token.T_2: true,
	},
	// AxBC : AorB ∙c
	{
		token.T_2: true,
	},
	// AxBC : AorB c ∙
	{
	},
}