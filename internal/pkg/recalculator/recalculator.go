package recalculator

import (
	"errors"
	"fmt"
	"github.com/mnogu/go-calculator"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/user"
	"log"
	"math"
	"strconv"
	"strings"
	"unsafe"
)

func ReCalculate(formulaArr []user.Formula, data *map[string]interface{}, tempJson []user.DataByFormID) error {
	for k, fr := range formulaArr {
		if k == len(formulaArr) {
			k--
		}
		if len(fr.TargetCells) != 2 {
			return errors.New("invalid target")
		}

		var (
			sum              float64
			currentRow       []int
			currentCol       int
			currentTargetRow []int
			currentTargetCol int
		)
		layout := fr.Layout
		dataMap := *(*map[string]interface{})(unsafe.Pointer(data))
		if dataMap["DATA"] == nil || dataMap["COLUMNS"] == nil {
			return errors.New("empty data")
		}

		dataMapData, ok := dataMap["DATA"].([]interface{})
		columnMapColumns, ok := dataMap["COLUMNS"].([]interface{})
		if !ok {
			return errors.New("invalid data format")
		}

		switch fr.TargetCells[0].(type) {
		case string:
			for n, _ := range dataMapData {
				if len(currentTargetRow) == 0 || currentTargetRow[n] != n+1 {
					currentTargetRow = append(currentTargetRow, n+1)
				}
			}
			currentTargetCol = fr.TargetCells[1].(int)
		case int:
			if len(currentTargetRow) < 1 {
				currentTargetRow = append(currentTargetRow, fr.TargetCells[0].(int))
			}
			currentTargetRow[0] = fr.TargetCells[0].(int)
			currentTargetCol = fr.TargetCells[1].(int)
		default:
			return errors.New("invalid type of target")
		}

		for c, _ := range currentTargetRow {
			for _, src := range fr.SourceCells {
				var (
					val interface{}
					err error
				)
				switch src[0].(type) {
				case string:
					for n, _ := range dataMapData {
						if len(currentRow) == 0 || currentRow[n] != n+1 {
							currentRow = append(currentRow, n+1)
						}
					}
					currentCol = src[1].(int)
				case int:
					if len(src) != 1 {
						if len(currentRow) < 1 {
							currentRow = append(currentRow, src[0].(int))
						}
						currentRow[0] = src[0].(int)
						currentCol = src[1].(int)
					} else {
						if len(currentRow) < 1 {
							currentRow = append(currentRow, src[0].(int))
						}
						currentRow[0] = src[0].(int)
					}
				default:
					return errors.New("invalid type of expression")
				}

				switch len(src) {
				case 1:
					val = strconv.Itoa(currentRow[0])
					currentRow = append(currentRow[:0], currentRow[1:]...)
				case 2:
					if currentRow[c] > len(dataMapData) || currentRow[c] < 1 {
						return errors.New("invalid row key")
					}
					sourceMap := dataMapData[currentRow[c]-1].(map[string]interface{})

					if currentCol > len(columnMapColumns) || currentCol < 1 {
						return errors.New("invalid column key")
					}
					columnMap := columnMapColumns[currentCol-1].(map[string]interface{})
					for _, x := range formulaArr {
						if x.TargetCells[0] == "#" {
							x.TargetCells[0] = currentRow[c]
						}
						if currentRow[c] == x.TargetCells[0].(int) && currentCol == x.TargetCells[1].(int) {
							var priorityFormula []user.Formula
							priorityFormula = append(priorityFormula, x)
							err = ReCalculate(priorityFormula, data, tempJson)
							if err != nil {
								log.Println("cannot calculate priority formula")
								//return errors.New("cannot calculate priority formula")
							}
						}
					}
					val, err = getMapValue(
						sourceMap,
						columnMap["VALUE_FIELD"].(string))
					if err != nil {
						return err
					}
					if val == nil {
						val = 0
					}
				case 3:
					for _, temp := range tempJson {
						if src[2] == temp.TemplateID {
							if temp.Data["DATA"] == nil || temp.Data["COLUMNS"] == nil {
								return errors.New("empty data")
							}

							tempMapData, ok := temp.Data["DATA"].([]interface{})
							tempColMapData, ok := temp.Data["COLUMNS"].([]interface{})
							if !ok {
								return errors.New("invalid data format")
							}

							if currentRow[c] > len(tempMapData) || currentRow[c] < 1 {
								return errors.New("invalid row key")
							}
							sourceMap := tempMapData[currentRow[c]-1].(map[string]interface{})

							if currentCol > len(tempColMapData) || currentCol < 1 {
								return errors.New("invalid column key")
							}
							columnMap := tempColMapData[currentCol-1].(map[string]interface{})

							val, err = getMapValue(
								sourceMap,
								columnMap["VALUE_FIELD"].(string))
							if err != nil {
								return err
							}
							if val == nil {
								val = 0
							}
						}
					}
				default:
					return errors.New("invalid val")
				}

				switch x := val.(type) {
				case string:
					strVal, err := strconv.Atoi(x)
					if err != nil {
						return fmt.Errorf("value cannot be string: %T", val)
					}
					newVal := strconv.Itoa(strVal)
					layout = strings.Replace(layout, "%s", newVal, 1)
				case float64:
					strVal := strconv.FormatFloat(x, 'f', 2, 32)
					layout = strings.Replace(layout, "%s", strVal, 1)
				case int:
					strVal := strconv.Itoa(x)
					layout = strings.Replace(layout, "%s", strVal, 1)
				}
			}

			s, err := calculator.Calculate(layout)
			if err != nil {
				sum = 0
				log.Println("failed calculate formula value")
				//return errors.New("failed calculate formula value")
			}
			if math.IsNaN(s) {
				return errors.New("division by zero is impossible")
			}
			strSum := fmt.Sprintf("%.3f", s)
			sum, err = strconv.ParseFloat(strSum, 64)
			if err != nil {
				return errors.New("failed conversion string to float")
			}

			targetMap, ok := dataMapData[currentTargetRow[c]-1].(map[string]interface{})
			if !ok {
				return errors.New("invalid json data")
			}

			columnMap, ok := columnMapColumns[currentTargetCol-1].(map[string]interface{})
			if !ok {
				return errors.New("invalid json data")
			}

			valBeforeFormula, err := getMapValue(targetMap, columnMap["VALUE_FIELD"].(string))
			if err != nil {
				return err
			}
			switch z := valBeforeFormula.(type) {
			case string:
				valBeforeFormula, err = strconv.ParseFloat(z, 64)
				if err != nil {
					return fmt.Errorf("value cannot be a string: %T", valBeforeFormula)
				}
			case nil:
				log.Println("no entered value")
				valBeforeFormula = 0.0
			}
			if sum != valBeforeFormula.(float64) {
				log.Println("formula value entered incorrectly")
				if err = setMapValue(&targetMap, columnMap["VALUE_FIELD"].(string), sum); err != nil {
					return err
				}
			}
			if err = setMapValue(&targetMap, columnMap["VALUE_FIELD"].(string), sum); err != nil {
				return err
			}
		}
	}
	return nil
}

func getMapValue(m map[string]interface{}, k string) (interface{}, error) {
	keys := strings.Split(k, ".")
	temp := m
	var idx int
	for i, key := range keys {
		if strings.Contains(key, "[") {
			path := strings.Split(strings.TrimRight(key, "]"), "[")
			key = path[0]
			idx, _ = strconv.Atoi(path[1])
		}
		switch x := temp[key].(type) {
		case map[string]interface{}:
			temp = x
		case []interface{}:
			switch y := x[idx].(type) {
			case map[string]interface{}:
				temp = y
			default:
				if i != len(keys)-1 {
					log.Printf("type: %T\nvalue: %v\n", x, x)
					log.Println("invalid reference in getting value")
					//return nil, errors.New("invalid reference")
				} else {
					return x, nil
				}
			}
		default:
			if i != len(keys)-1 {
				log.Printf("type: %T\nvalue: %v\n", x, x)
				log.Println("invalid reference in getting value")
				//return nil, errors.New("invalid reference")
			} else {
				return x, nil
			}
		}
	}
	return nil, errors.New("error in getting data value")
}

func setMapValue(m *map[string]interface{}, k string, v interface{}) error {
	keys := strings.Split(k, ".")
	temp := *m
	var idx int
	for i, key := range keys {
		if strings.Contains(key, "[") {
			path := strings.Split(strings.TrimRight(key, "]"), "[")
			key = path[0]
			idx, _ = strconv.Atoi(path[1])
		}
		switch x := temp[key].(type) {
		case map[string]interface{}:
			temp = x
		case []interface{}:
			switch y := x[idx].(type) {
			case map[string]interface{}:
				temp = y
			default:
				if i != len(keys)-1 {
					log.Printf("type: %T\nvalue: %v\n", x, x)
					log.Println("invalid reference in setting value")
					//return errors.New("invalid reference")
				} else {
					temp[keys[i]] = v
				}
			}
		default:
			if i != len(keys)-1 {
				log.Printf("type: %T\nvalue: %v\n", x, x)
				log.Println("invalid reference in setting value")
				//return errors.New("invalid reference")
			} else {
				temp[keys[i]] = v
			}
		}
	}
	return nil
}
