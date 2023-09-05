package auth

import (
	"errors"
	"github.com/Knetic/govaluate"
	"time"
)

func inBetweenDateFinish() govaluate.ExpressionFunction {
	return func(args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return false, errors.New("недостаточно парамаетров для проверки inBetweenDateFinish")
		}

		switch date1 := args[0].(type) {
		case string:
			obj, err := time.Parse(layout, date1)
			if err != nil {
				return false, errors.New("неверна введена дата finish")
			}
			switch date2 := args[1].(type) {
			case string:
				sub, err := time.Parse(layout, date2)
				if err != nil {
					return false, errors.New("неверна введена дата finish")
				}
				result := sub.After(obj)
				if !result {
					return false, errors.New("дата не соответсвует доступной finish")
				} else {
					return true, err
				}
			default:
				return false, errors.New("невозможно обработать дату finish")

			}

		default:
			return false, errors.New("первый аргумент не дата finish")
		}
	}
}
