package auth

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	"time"
)

const layout = "2006-01-02"

func inBetweenDate() govaluate.ExpressionFunction {
	return func(args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return false, errors.New("недостаточно парамаетров для проверки inBetweenDate")
		}

		switch date1 := args[0].(type) {
		case string:
			obj, err := time.Parse(layout, date1)
			if err != nil {
				return false, errors.New("неверна введена дата")
			}
			switch date2 := args[1].(type) {
			case string:
				sub, err := time.Parse(layout, date2)
				if err != nil {
					return false, errors.New("неверна введена дата")
				}
				result := sub.Before(obj)
				if !result {
					fmt.Println(sub, obj)
					return false, errors.New("дата не соответсвует доступной")
				} else {
					return true, err
				}
			default:
				return false, errors.New("невозможно обработать дату")

			}

		default:
			return false, errors.New("первый аргумент не дата")
		}
	}
}
