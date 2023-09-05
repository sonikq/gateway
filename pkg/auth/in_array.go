package auth

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
)

func inArray() govaluate.ExpressionFunction {
	return func(args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return false, errors.New("недостаточно парамаетров для проверки inArray")
		}

		switch arr := args[0].(type) {
		case []int:
			if len(arr) == 0 {
				return false, errors.New("пустой массив")
			}

			switch item := args[1].(type) {
			case int:
				for _, arrItem := range arr {
					if arrItem == item {
						return true, nil
					}
				}
				return false, nil
			case float64:
				for _, arrItem := range arr {
					if arrItem == int(item) {
						return true, nil
					}
				}
				return false, nil
			default:
				return false, errors.New("тип второго аргумента не сходится с массивом")

			}
		case []int64:
			if len(arr) == 0 {
				return false, errors.New("пустой массив")
			}

			switch item := args[1].(type) {
			case int64:
				for _, arrItem := range arr {
					if arrItem == item {
						return true, nil
					}
				}
				return false, nil
			case float64:
				for _, arrItem := range arr {
					if arrItem == int64(item) {
						return true, nil
					}
				}
				return false, nil
			default:
				return false, errors.New("тип второго аргумента не сходится с массивом")

			}

		case []string:
			if len(arr) == 0 {
				return false, errors.New("пустой массив")
			}

			switch item := args[1].(type) {
			case string:
				for _, arrItem := range arr {
					if arrItem == item {
						return true, nil
					}
				}
				fmt.Println(args)
				return false, nil
			default:
				return false, errors.New("тип второго аргумента не сходится с массивом")

			}

		default:
			fmt.Printf("type is %T", arr)
			return false, errors.New("первый аргумент не массив")
		}
	}
}
