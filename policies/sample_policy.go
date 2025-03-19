package policies

import "github.com/OdairPianta/julia/models"

type SamplePolicy struct {
	User *models.User
}

func NewSamplePolicy(user *models.User) *SamplePolicy {
	return &SamplePolicy{User: user}
}

func (p *SamplePolicy) ViewAny() bool {
	if p.User == nil {
		return false
	}
	return true
}

func (p *SamplePolicy) View(model *models.SampleModel) bool {
	if p.User == nil || model == nil {
		return false
	}
	return true
}

func (p *SamplePolicy) Create() bool {
	if p.User == nil {
		return false
	}
	return true
}

func (p *SamplePolicy) Delete(model *models.SampleModel) bool {
	if p.User == nil || model == nil {
		return false
	}
	return true
}

func (p *SamplePolicy) Update(model *models.SampleModel) bool {
	if p.User == nil || model == nil {
		return false
	}
	return true
}
