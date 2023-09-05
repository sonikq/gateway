package auth

import (
	"github.com/casbin/casbin/v2"
)

func NewEnforcer(modelPath, policyPath string) (*casbin.Enforcer, error) {
	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		return nil, err
	}

	enforcer.AddFunction("inArray", inArray())
	enforcer.AddFunction("inBetweenDate", inBetweenDate())
	enforcer.AddFunction("inBetweenDateFinish", inBetweenDateFinish())
	return enforcer, nil
}
