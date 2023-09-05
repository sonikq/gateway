package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Formula struct {
	// Target example: [1][4] => {1, 4}
	Target []interface{}

	// Source
	// example 1: [1][4] => {1, 4}
	// example 2: [1][4][1233145] => {1, 4, 1233145},
	// where 1233145 is chapter_template_id
	Source [][]interface{}

	// Layout example: [1][4]+[1][6] => "=%s+%s"
	Layout string
}

func ParseFormula(target, expression string) (*Formula, error) {
	targetCells, err := parseTarget(target)
	if err != nil {
		return nil, err
	}

	sourceCells, layout, err := parseExpression(expression)
	if err != nil {
		return nil, err
	}

	return &Formula{
		Target: targetCells,
		Source: sourceCells,
		Layout: layout,
	}, nil
}

func parseTarget(s string) ([]interface{}, error) {
	indexes, err := parseIndexes(s)
	if err != nil {
		return nil, err
	}

	if len(indexes) != 2 {
		errMsg := fmt.Sprintf("invalid target: %s, please check target cell is valid!", s)
		return nil, errors.New(errMsg)
	}

	return indexes, nil
}

func parseExpression(s string) ([][]interface{}, string, error) {
	var arr [][]interface{}

	operands := getOperands(s)

	layout := s

	for _, operand := range operands {
		if !strings.Contains(operand, "[") && !strings.Contains(operand, "]") {
			num, err := strconv.Atoi(operand)
			if err != nil {
				return nil, "", err
			}
			arr = append(arr, []interface{}{num})
		} else {
			idx, err := parseIndexes(operand)
			if err != nil {
				return nil, "", err
			}
			arr = append(arr, idx)
		}

		layout = strings.Replace(layout, operand, "%s", 1)
	}

	//if strings.Count(layout, "%s")*2 != len(arr) {
	//	return nil, "", errors.New("formula parser is not working properly")
	//}

	return arr, layout, nil
}

func parseIndexes(s string) ([]interface{}, error) {
	var (
		arr      []interface{}
		beg, end int
		opened   bool
	)

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[':
			if !opened {
				opened = true
				beg, end = i, -1
			} else {
				errMsg := fmt.Sprintf("invalid expression: %s, please check expression is valid!", s)
				return nil, errors.New(errMsg)
			}
		case ']':
			if opened {
				opened = false
				end = i
			} else {
				errMsg := fmt.Sprintf("invalid expression: %s, please check expression is valid!", s)
				return nil, errors.New(errMsg)
			}
		}

		if end > beg {
			if s[beg+1:end] == "#" {
				arr = append(arr, s[beg+1:end])
			} else {
				num, err := strconv.Atoi(s[beg+1 : end])
				if err != nil {
					return nil, err
				}
				arr = append(arr, num)
			}
			beg, end = -1, -1
		}
	}
	return arr, nil
}

func getOperands(s string) []string {
	r := strings.NewReplacer(
		"(", "",
		")", "",
		"+", " ",
		"-", " ",
		"*", " ",
		"/", " ")

	operands := strings.Split(r.Replace(s), " ")
	return operands
}

//func getOperands(s string) []string {
//	var indexes []string
//	beg, end := 0, -1
//	for i := 0; i < len(s); i++ {
//		switch s[i] {
//		case '+', '-', '*', '/', ')':
//			end = i
//			indexes = append(indexes, s[beg:end])
//			beg = i + 1
//		case '(':
//			beg = i + 1
//		default:
//			if i == len(s)-1 {
//				end = i
//				indexes = append(indexes, s[beg:end])
//			}
//		}
//	}
//	return indexes
//}
