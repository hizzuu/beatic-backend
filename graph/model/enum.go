package model

import (
	"errors"

	"github.com/hizzuu/beatic-backend/internal/domain"
)

var gender = map[string]domain.Gender{
	GenderMale.String():      domain.GenderMale,
	GenderFemale.String():    domain.GenderFemale,
	GenderNonbinary.String(): domain.GenderNonbinary,
	GenderOther.String():     domain.GenderOther,
	GenderNoanswer.String():  domain.GenderNoanswer,
}

func (t Gender) ConvGender() domain.Gender {
	return gender[t.String()]
}

func ConvGender(g domain.Gender) (Gender, error) {
	switch g {
	case domain.GenderMale:
		return GenderMale, nil
	case domain.GenderFemale:
		return GenderFemale, nil
	case domain.GenderNonbinary:
		return GenderNonbinary, nil
	case domain.GenderOther:
		return GenderOther, nil
	case domain.GenderNoanswer:
		return GenderNoanswer, nil
	default:
		return Gender(""), errors.New("unexpected domain.Gender type")
	}
}
