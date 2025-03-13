package policies

import (
	"github.com/OdairPianta/julia/enums"
	"github.com/OdairPianta/julia/models"
)

type UserPolicy struct {
	User *models.User
}

func NewUserPolicy(user *models.User) *UserPolicy {
	return &UserPolicy{User: user}
}

func (p *UserPolicy) ViewAny() bool {
	if p.User == nil {
		return false
	}
	return p.User.Profile == enums.UserProfileEnumAdministrator
}

func (p *UserPolicy) View(model *models.User) bool {
	if p.User == nil || model == nil {
		return false
	}
	return p.User.ID == model.ID || p.User.Profile == enums.UserProfileEnumAdministrator
}

func (p *UserPolicy) Create() bool {
	if p.User == nil {
		return false
	}
	return p.User.Profile == enums.UserProfileEnumAdministrator
}

func (p *UserPolicy) Delete(model *models.User) bool {
	if p.User == nil || model == nil {
		return false
	}
	return p.User.ID == model.ID || p.User.Profile == enums.UserProfileEnumAdministrator
}

func (p *UserPolicy) Update(model *models.User) bool {
	if p.User == nil || model == nil {
		return false
	}
	return p.User.ID == model.ID || p.User.Profile == enums.UserProfileEnumAdministrator
}
