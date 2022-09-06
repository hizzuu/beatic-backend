package domain

type Gender int64

const (
	_ Gender = iota
	GenderMale
	GenderFemale
	GenderNonbinary
	GenderOther
	GenderNoanswer
)
