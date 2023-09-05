package recalculator

//
//import (
//	"encoding/json"
//	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models/user"
//	"testing"
//)
//
//func TestReCalculate(t *testing.T) {
//	f := []user.Formula{
//
//	}
//
//	var data user.SaveChapterData
//	if err := json.Unmarshal(s, &data); err != nil {
//		t.Error(err)
//	}
//
//	if err := ReCalculate(fMap, &data); err != nil {
//		t.Error(err)
//	}
//}
//
//var s = []byte(`
//{
//    "ID": 100091,
//    "TITLE": "СВЕДЕНИЯ ОБ АВТОМОБИЛЬНЫХ ДОРОГАХ ОБЩЕГО ПОЛЬЗОВАНИЯ И СООРУЖЕНИЯХ НА НИХ ФЕДЕРАЛЬНОГО, РЕГИОНАЛЬНОГО ИЛИ МЕЖМУНИЦИПАЛЬНОГО ЗНАЧЕНИЯ",
//    "OKUD": "0615064",
//    "CODE": "0615064/1",
//    "LABEL": "Раздел 1",
//    "CUSTOM": false,
//    "CHAPTER": {
//        "TEXT":"Сведения об автомобильных дорогах общего пользования и сооружениях на них федерального, регионального, или межмуниципального значения",
//        "STYLE": {
//			"FONT": "Arial",
//			"SIZE": 10,
//			"FONT_STYLE": {
//				"BOLD": true
//			},
//			"VERT_ALIGNMENT": "Center",
//			"HOR_ALIGNMENT": "Center"
//		}
//    },
//    "PERIOD": "20230100",
//    "ORG_RESPONDER_ID": 3288127,
//    "ORG_RESPONDER": "КУ ВО \"Управление автомобильных дорог Вологодской области \"",
//    "ORG_REVIEWER_ID": 2323,
//    "EXECUTOR_ID": 117754818,
//    "EXECUTOR_FIO": "Кормакова А.Р.",
//    "REGION": "Вологодская область",
//    "REGION_OKATO": "19",
//    "REGION_ID": 24595,
//    "PRE_HEADER": {
//        "TEXT": "",
//        "STYLE": {
//            "FONT": "Arial",
//            "SIZE": 8,
//            "FONT_STYLE": {"BOLD":true},
//            "VERT_ALIGNMENT": "Center",
//            "HOR_ALIGNMENT": "Right"
//        }
//    },
//    "HEADER": {
//        "HEADER_COLUMNS": 13,
//        "HEADER_ROWS": 4,
//        "STYLE": {
//            "FONT": "Arial",
//            "SIZE": 10,
//            "VERT_ALIGNMENT": "Center",
//            "HOR_ALIGNMENT": "Center",
//            "FONT_STYLE": {
//                "BOLD": true
//            }
//        },
//        "HEADERS": [
//            {
//                "HEADER_ID": 1,
//                "HEADER_CAPTION": "Наименование показателя",
//                "HEADER_GRID": [
//                    1,
//                    1,
//                    4,
//                    1
//                ]
//            },
//            {
//                "HEADER_ID": 2,
//                "HEADER_CAPTION": "№ строки",
//                "HEADER_GRID": [
//                    1,
//                    2,
//                    4,
//                    2
//                ]
//            },
//            {
//                "HEADER_ID": 3,
//                "HEADER_CAPTION": "Единица измерения",
//                "HEADER_GRID": [
//                    1,
//                    3,
//                    4,
//                    3
//                ]
//            },
//            {
//                "HEADER_ID": 4,
//                "HEADER_CAPTION": "Наличие на начало \nотчетного года",
//                "HEADER_GRID": [
//                    1,
//                    4,
//                    4,
//                    4
//                ]
//            },
//            {
//                "HEADER_ID": 5,
//                "HEADER_CAPTION": "Движение за отчетный год",
//                "HEADER_GRID": [
//                    1,
//                    5,
//                    1,
//                    12
//                ]
//            },
//            {
//                "HEADER_ID": 6,
//                "HEADER_CAPTION": "Наличие на конец \nотчетного года",
//                "HEADER_GRID": [
//                    1,
//                    13,
//                    4,
//                    13
//                ]
//            },
//            {
//                "HEADER_ID": 7,
//                "HEADER_CAPTION": "увеличение",
//                "HEADER_GRID": [
//                    2,
//                    5,
//                    2,
//                    8
//                ]
//            },
//            {
//                "HEADER_ID": 8,
//                "HEADER_CAPTION": "уменьшение",
//                "HEADER_GRID": [
//                    2,
//                    9,
//                    2,
//                    12
//                ]
//            },
//            {
//                "HEADER_ID": 9,
//                "HEADER_CAPTION": "принято в эксплуатацию",
//                "HEADER_GRID": [
//                    3,
//                    5,
//                    3,
//                    6
//                ]
//            },
//            {
//                "HEADER_ID": 10,
//                "HEADER_CAPTION": "включено или принято \nв данный перечень дорог",
//                "HEADER_GRID": [
//                    3,
//                    7,
//                    3,
//                    8
//                ]
//            },
//            {
//                "HEADER_ID": 11,
//                "HEADER_CAPTION": "после строительства и реконструкции",
//                "HEADER_GRID": [
//                    4,
//                    5,
//                    4,
//                    5
//                ]
//            },
//            {
//                "HEADER_ID": 12,
//                "HEADER_CAPTION": "после капитального ремонта и ремонта",
//                "HEADER_GRID": [
//                    4,
//                    6,
//                    4,
//                    6
//                ]
//            },
//            {
//                "HEADER_ID": 13,
//                "HEADER_CAPTION": "после строительства \nи реконструкции",
//                "HEADER_GRID": [
//                    3,
//                    9,
//                    4,
//                    9
//                ]
//            },
//            {
//                "HEADER_ID": 14,
//                "HEADER_CAPTION": "после капитального ремонта и ремонта",
//                "HEADER_GRID": [
//                    3,
//                    10,
//                    4,
//                    10
//                ]
//            },
//            {
//                "HEADER_ID": 15,
//                "HEADER_CAPTION": "переведено в другой перечень \nдорог или передано другим организациям",
//                "HEADER_GRID": [
//                    3,
//                    11,
//                    4,
//                    11
//                ]
//            },
//            {
//                "HEADER_ID": 16,
//                "HEADER_CAPTION": "по итогам инвентаризации, паспортизации, списано",
//                "HEADER_GRID": [
//                    3,
//                    12,
//                    4,
//                    12
//                ]
//            },
//            {
//                "HEADER_ID": 17,
//                "HEADER_CAPTION": "из другого перечня \nдорог или от других \nорганизаций",
//                "HEADER_GRID": [
//                    4,
//                    7,
//                    4,
//                    7
//                ]
//            },
//            {
//                "HEADER_ID": 18,
//                "HEADER_CAPTION": "по итогам инвентаризации, паспортизации, списано",
//                "HEADER_GRID": [
//                    4,
//                    8,
//                    4,
//                    8
//                ]
//            }
//        ]
//    },
//    "COMMON": {
//        "OKEI_COLUMN": 3,
//        "ROW_NUM_COLUMN": 2,
//        "ROW_CAPTION_COLUMN": 1,
//        "ROW_UNIT_WIDTH": 17,
//        "STATIC_ROWS": true,
//        "STYLE": {
//            "FONT": "Arial",
//            "HOR_ALIGNMENT": "Center",
//            "VERT_ALIGNMENT": "Center",
//            "SIZE": 8
//        }
//    },
//    "COLUMNS": [
//        {
//            "COLUMN_ID": 1,
//            "COLUMN_NUM": "1",
//            "COLUMN_WIDTH_PX": 40,
//            "COLUMN_WIDTH_MM": 65,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left",
//                "VERT_ALIGNMENT": "Center"
//            },
//            "DATA_TYPE": "TEXT",
//            "VALUE_FIELD": "PARAM_NAME",
//            "IS_EDITABLE": false
//        },
//        {
//            "COLUMN_ID": 2,
//            "COLUMN_NUM": "2",
//            "COLUMN_WIDTH_PX": 10,
//            "COLUMN_WIDTH_MM": 10,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Center",
//                "VERT_ALIGNMENT": "Center"
//            },
//            "DATA_TYPE": "INT",
//            "VALUE_FIELD": "ROW_NUM",
//            "IS_EDITABLE": false
//        },
//        {
//            "COLUMN_ID": 3,
//            "COLUMN_NUM": "3",
//            "COLUMN_WIDTH_PX": 10,
//            "COLUMN_WIDTH_MM": 12,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Center",
//                "VERT_ALIGNMENT": "Center"
//            },
//            "DATA_TYPE": "SELECTOR",
//            "VALUE_FIELD": "OKEI_ID",
//            "VALUE_LABEL": "OKEI",
//            "SELECTOR": {
//                "DATA_SOURCE": "OKEI",
//                "VALUE_FIELD": "ID",
//                "VALUE_LABEL": "NAME"
//            },
//            "IS_EDITABLE": false
//        },
//        {
//            "COLUMN_ID": 4,
//            "COLUMN_NUM": "4",
//            "COLUMN_WIDTH_PX": 10,
//            "COLUMN_WIDTH_MM": 16,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Right",
//                "VERT_ALIGNMENT": "Center"
//            },
//            "DATA_TYPE": "FROM_OKEI",
//            "VALUE_FIELD": "VALUE.BEG"
//        },
//        {
//            "COLUMN_ID": 5,
//            "COLUMN_NUM": "5",
//            "COLUMN_WIDTH_PX": 10,
//            "COLUMN_WIDTH_MM": 20,
//            "VALUE_FIELD": "VALUE.CONSTRUCTED.PLUS",
//            "DATA_TYPE": "FROM_OKEI",
//            "STYLE": {
//                "HOR_ALIGNMENT": "Right",
//                "VERT_ALIGNMENT": "Center"
//            }
//        },
//        {
//            "COLUMN_ID": 6,
//            "COLUMN_NUM": "6",
//            "COLUMN_WIDTH_PX": 10,
//            "COLUMN_WIDTH_MM": 20,
//            "VALUE_FIELD": "VALUE.REPAIRED.PLUS",
//            "DATA_TYPE": "FROM_OKEI",
//            "STYLE": {
//                "HOR_ALIGNMENT": "Right",
//                "VERT_ALIGNMENT": "Center"
//            }
//        },
//        {
//            "COLUMN_ID": 7,
//            "COLUMN_NUM": "7",
//            "COLUMN_WIDTH_PX": 10,
//            "COLUMN_WIDTH_MM": 20,
//            "VALUE_FIELD": "VALUE.LISTED.PLUS",
//            "DATA_TYPE": "FROM_OKEI",
//            "STYLE": {
//                "HOR_ALIGNMENT": "Right",
//                "VERT_ALIGNMENT": "Center"
//            }
//        },
//        {
//            "COLUMN_ID": 8,
//            "COLUMN_NUM": "8",
//            "COLUMN_WIDTH_PX": 10,
//            "COLUMN_WIDTH_MM": 20,
//            "VALUE_FIELD": "VALUE.INVENTORIED.PLUS",
//            "DATA_TYPE": "FROM_OKEI",
//            "STYLE": {
//                "HOR_ALIGNMENT": "Right",
//                "VERT_ALIGNMENT": "Center"
//            }
//        },
//        {
//            "COLUMN_ID": 9,
//            "COLUMN_NUM": "9",
//            "COLUMN_WIDTH_PX": 10,
//            "COLUMN_WIDTH_MM": 20,
//            "VALUE_FIELD": "VALUE.CONSTRUCTED.MINUS",
//            "DATA_TYPE": "FROM_OKEI",
//            "STYLE": {
//                "HOR_ALIGNMENT": "Right",
//                "VERT_ALIGNMENT": "Center"
//            }
//        },
//        {
//            "COLUMN_ID": 10,
//            "COLUMN_NUM": "10",
//            "COLUMN_WIDTH_PX": 10,
//            "COLUMN_WIDTH_MM": 20,
//            "VALUE_FIELD": "VALUE.REPAIRED.MINUS",
//            "DATA_TYPE": "FROM_OKEI",
//            "STYLE": {
//                "HOR_ALIGNMENT": "Right",
//                "VERT_ALIGNMENT": "Center"
//            }
//        },
//        {
//            "COLUMN_ID": 11,
//            "COLUMN_NUM": "11",
//            "COLUMN_WIDTH_PX": 10,
//            "COLUMN_WIDTH_MM": 20,
//            "VALUE_FIELD": "VALUE.LISTED.MINUS",
//            "DATA_TYPE": "FROM_OKEI",
//            "STYLE": {
//                "HOR_ALIGNMENT": "Right",
//                "VERT_ALIGNMENT": "Center"
//            }
//        },
//        {
//            "COLUMN_ID": 12,
//            "COLUMN_NUM": "12",
//            "COLUMN_WIDTH_PX": 10,
//            "COLUMN_WIDTH_MM": 20,
//            "VALUE_FIELD": "VALUE.INVENTORIED.MINUS",
//            "DATA_TYPE": "FROM_OKEI",
//            "STYLE": {
//                "HOR_ALIGNMENT": "Right",
//                "VERT_ALIGNMENT": "Center"
//            }
//        },
//        {
//            "COLUMN_ID": 13,
//            "COLUMN_NUM": "13",
//            "COLUMN_WIDTH_PX": 10,
//            "COLUMN_WIDTH_MM": 20,
//            "VALUE_FIELD": "VALUE.END",
//            "DATA_TYPE": "FROM_OKEI",
//            "STYLE": {
//                "HOR_ALIGNMENT": "Right",
//                "VERT_ALIGNMENT": "Center"
//            }
//        }
//    ],
//    "ROW_UNIONS": [
//        {
//            "GRID":[31,1,32,1]
//        },
//        {
//            "GRID":[35,1,36,1]
//        },
//        {
//            "GRID":[39,1,40,1]
//        },
//        {
//            "GRID":[41,1,42,1]
//        },
//        {
//            "GRID":[47,1,48,1]
//        },
//        {
//            "GRID":[50,1,51,1]
//        },
//        {
//            "GRID":[53,1,54,1]
//        },
//        {
//            "GRID":[55,1,56,1]
//        },
//        {
//            "GRID":[57,1,58,1]
//        },
//        {
//            "GRID":[59,1,60,1]
//        },
//        {
//            "GRID":[61,1,62,1]
//        },
//        {
//            "GRID":[63,1,64,1]
//        },
//        {
//            "GRID":[65,1,66,1]
//        },
//        {
//            "GRID":[67,1,68,1]
//        },
//        {
//            "GRID":[71,1,72,1]
//        },
//        {
//            "GRID":[75,1,76,1]
//        },
//        {
//            "GRID":[77,1,78,1]
//        },
//        {
//            "GRID":[79,1,80,1]
//        }
//    ],
//    "ROW_BREAKS": [
//        {
//            "PREV_ROW": 0,
//            "BREAK_CAPTION": "1. АВТОМОБИЛЬНЫЕ ДОРОГИ",
//            "INDENT": 0,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left",
//                "FONT_STYLE": {
//                    "BOLD": true
//                }
//            }
//        },
//        {
//            "PREV_ROW": 1,
//            "BREAK_CAPTION": "в том числе:",
//            "INDENT": 0,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left"
//            }
//        },
//        {
//            "PREV_ROW": 3,
//            "BREAK_CAPTION": "из них:",
//            "INDENT": 4,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left"
//            }
//        },
//        {
//            "PREV_ROW": 7,
//            "BREAK_CAPTION": "из них:",
//            "INDENT": 4,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left"
//            }
//        },
//        {
//            "PREV_ROW": 11,
//            "BREAK_CAPTION": "в том числе:",
//            "INDENT": 0,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left"
//            }
//        },
//        {
//            "PREV_ROW": 12,
//            "BREAK_CAPTION": "из них",
//            "INDENT": 2,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left"
//            }
//        },
//        {
//            "PREV_ROW": 20,
//            "BREAK_CAPTION": "2. ДОРОЖНЫЕ СООРУЖЕНИЯ \n 2.1. Мосты, путепроводы и эстакады - всего,",
//            "INDENT": 0,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left",
//                "FONT_STYLE": {
//                    "BOLD": true
//                }
//            }
//        },
//        {
//            "PREV_ROW": 22,
//            "BREAK_CAPTION": "в том числе капитальные - всего",
//            "INDENT": 2,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left"
//            }
//        },
//        {
//            "PREV_ROW": 26,
//            "BREAK_CAPTION": "из капитальных мостов и путепроводов:",
//            "INDENT": 2,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left"
//            }
//        },
//        {
//            "PREV_ROW": 44,
//            "BREAK_CAPTION": "2.3. Пешеходные переходы в разных \nуровнях - всего",
//            "INDENT": 0,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left",
//                "FONT_STYLE":{"BOLD":true}
//            }
//        },
//        {
//            "PREV_ROW": 46,
//            "BREAK_CAPTION": " в том числе:",
//            "INDENT": 0,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left"
//            }
//        },
//        {
//            "PREV_ROW": 54,
//            "BREAK_CAPTION": " в том числе:",
//            "INDENT": 2,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left"
//            }
//        },
//        {
//            "PREV_ROW": 70,
//            "BREAK_CAPTION": " в том числе:",
//            "INDENT": 0,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left"
//            }
//        },
//        {
//            "PREV_ROW": 83,
//            "BREAK_CAPTION": "из них:",
//            "INDENT": 0,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left"
//            }
//        },
//        {
//            "PREV_ROW": 84,
//            "BREAK_CAPTION": "в том числе:",
//            "INDENT": 2,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left"
//            }
//        },
//        {
//            "PREV_ROW": 94,
//            "BREAK_CAPTION": "Справочно: \nиз строки 01 по графе 13:",
//            "INDENT": 0,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left",
//                "FONT_STYLE":{"BOLD":true}
//            }
//        },
//        {
//            "PREV_ROW": 97,
//            "BREAK_CAPTION": "из строки 23 по графе 13:",
//            "INDENT": 0,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left",
//                "FONT_STYLE":{"BOLD":true}
//            }
//        },
//        {
//            "PREV_ROW": 99,
//            "BREAK_CAPTION": "из строки 24 по графе 13:",
//            "INDENT": 0,
//            "STYLE": {
//                "HOR_ALIGNMENT": "Left",
//                "FONT_STYLE":{"BOLD":true}
//            }
//        }
//    ],
//    "READONLY": [
//        [
//            83,
//            5,
//            94,
//            13
//        ]
//    ],
//    "DATA": [
//        {
//            "PARAM_NAME": "1.1. Общая протяженность дорог - всего\n(сумма строк 02 и 10)",
//            "ROW_NUM": 1,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "BEG": 11,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 123
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "с твердым покрытием,\nсумма строк 03 и 07)",
//            "ROW_NUM": 2,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "BEG": 22,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "из строки 02 - с усовершенствованным \n покрытием (сумма строк 04 ÷ 06)",
//            "ROW_NUM": 3,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "цементобетонные",
//            "ROW_NUM": 4,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "асфальтобетонные",
//            "ROW_NUM": 5,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "из щебня и гравия, обработанных \nвяжущими материалами",
//            "ROW_NUM": 6,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "из строки 02 - с покрытием\nпереходного типа \n(сумма строк 08 и 09)",
//            "ROW_NUM": 7,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "из щебня и гравия (шлака),\n не обработанных вяжущими материалами, каменные мостовые",
//            "ROW_NUM": 8,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "из грунтов и местных малопрочных \nматериалов, обработанных вяжущими \nматериалами",
//            "ROW_NUM": 9,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "грунтовые",
//            "ROW_NUM": 10,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "BEG": 33,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "1.1.1. Общая площадь твердых типов \nпокрытия дорог - всего \n(сумма строк 12 и 16)",
//            "ROW_NUM": 11,
//            "OKEI_ID": 58,
//            "OKEI": "тыс м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "площадь усовершенствованных типов \nпокрытия(сумма строк 13 ÷ 15)",
//            "ROW_NUM": 12,
//            "OKEI_ID": 58,
//            "OKEI": "тыс м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "цементобетонных покрытий\n(производная строки 04)",
//            "ROW_NUM": 13,
//            "OKEI_ID": 58,
//            "OKEI": "тыс м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "асфальтобетонных покрытий \n(производная строки 05)",
//            "ROW_NUM": 14,
//            "OKEI_ID": 58,
//            "OKEI": "тыс м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "из щебня и гравия, обработанных \nвяжущими материалами \n(производная строки 06)",
//            "ROW_NUM": 15,
//            "OKEI_ID": 58,
//            "OKEI": "тыс м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "площадь покрытий переходного типа \n(производная строки 07)",
//            "ROW_NUM": 16,
//            "OKEI_ID": 58,
//            "OKEI": "тыс м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "1.2. Паромные переправы",
//            "ROW_NUM": 17,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "в том числе с применением самоходных плавсредств \n(самоходные баржи, буксиры)",
//            "ROW_NUM": 18,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "1.3. Автозимники и ледовые переправы",
//            "ROW_NUM": 19,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "в том числе ледовые переправы",
//            "ROW_NUM": 20,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "(сумма строк 23 и 39)",
//            "ROW_NUM": 21,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "(сумма строк 24 и 40)",
//            "ROW_NUM": 22,
//            "OKEI_ID": 18,
//            "OKEI": "пог.м",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "(сумма строк 27, 31, 35)",
//            "ROW_NUM": 23,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "(сумма строк 28, 32, 36)",
//            "ROW_NUM": 24,
//            "OKEI_ID": 18,
//            "OKEI": "пог.м",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "общая площадь капитальных мостов, \nпутепроводов и эстакад \n(сумма строк 29, 33, 37)",
//            "ROW_NUM": 25,
//            "OKEI_ID": 58,
//            "OKEI": "тыс. м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "в том числе площадь проезжей части \n(сумма строк 30, 34, 38)",
//            "ROW_NUM": 26,
//            "OKEI_ID": 58,
//            "OKEI": "тыс. м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "мосты железобетонные и \nкаменные",
//            "ROW_NUM": 27,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "мосты железобетонные и \nкаменные",
//            "ROW_NUM": 28,
//            "OKEI_ID": 18,
//            "OKEI": "пог.м",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "площадь мостов железобетонных \nи каменных",
//            "ROW_NUM": 29,
//            "OKEI_ID": 58,
//            "OKEI": "тыс. м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "в том числе площадь проезжей части",
//            "ROW_NUM": 30,
//            "OKEI_ID": 58,
//            "OKEI": "тыс. м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "мосты металлические",
//            "ROW_NUM": 31,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "мосты металлические",
//            "ROW_NUM": 32,
//            "OKEI_ID": 58,
//            "OKEI": "пог. м",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "площадь мостов металлических",
//            "ROW_NUM": 33,
//            "OKEI_ID": 58,
//            "OKEI": "тыс. м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "в том числе площадь проезжей \nчасти",
//            "ROW_NUM": 34,
//            "OKEI_ID": 58,
//            "OKEI": "тыс. м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "путепроводы и эстакады на \nпересечениях автомобильных дорог \nи с железными дорогами",
//            "ROW_NUM": 35,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "путепроводы и эстакады на \nпересечениях автомобильных дорог \nи с железными дорогами",
//            "ROW_NUM": 36,
//            "OKEI_ID": 58,
//            "OKEI": "пог. м",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "площадь путепроводов и эстакад",
//            "ROW_NUM": 37,
//            "OKEI_ID": 58,
//            "OKEI": "тыс. м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "  в том числе площадь проезжей части",
//            "ROW_NUM": 38,
//            "OKEI_ID": 58,
//            "OKEI": "тыс. м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "деревянные мосты",
//            "ROW_NUM": 39,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "деревянные мосты",
//            "ROW_NUM": 40,
//            "OKEI_ID": 58,
//            "OKEI": "пог. м",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "2.2 Тоннели автодорожные",
//            "ROW_NUM": 41,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "2.2 Тоннели автодорожные",
//            "ROW_NUM": 42,
//            "OKEI_ID": 58,
//            "OKEI": "пог. м",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "общая площадь проезжей части и \nпешеходных дорожек в тоннелях \nавтодорожных",
//            "ROW_NUM": 43,
//            "OKEI_ID": 58,
//            "OKEI": "тыс. м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "в том числе площадь проезжей части \nв тоннелях автодорожных",
//            "ROW_NUM": 44,
//            "OKEI_ID": 58,
//            "OKEI": "тыс. м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "(сумма строк 47 и 50)",
//            "ROW_NUM": 45,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "(сумма строк 48 и 51)",
//            "ROW_NUM": 46,
//            "OKEI_ID": 58,
//            "OKEI": "пог. м",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "эстакадные",
//            "ROW_NUM": 47,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "эстакадные",
//            "ROW_NUM": 48,
//            "OKEI_ID": 58,
//            "OKEI": "пог. м",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "общая площадь эстакадных пешеходных переходов",
//            "ROW_NUM": 49,
//            "OKEI_ID": 58,
//            "OKEI": "тыс. м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "подземные",
//            "ROW_NUM": 50,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "подземные",
//            "ROW_NUM": 51,
//            "OKEI_ID": 51,
//            "OKEI": "пог. м",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "площадь подземных пешеходных переходов",
//            "ROW_NUM": 52,
//            "OKEI_ID": 58,
//            "OKEI": "тыс. м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "2.4 Трубы - всего",
//            "ROW_NUM": 53,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "2.4 Трубы - всего",
//            "ROW_NUM": 54,
//            "OKEI_ID": 51,
//            "OKEI": "пог. м",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "железобетонные, бетонные и каменные",
//            "ROW_NUM": 55,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "железобетонные, бетонные и каменные",
//            "ROW_NUM": 56,
//            "OKEI_ID": 51,
//            "OKEI": "пог. м",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "металлические",
//            "ROW_NUM": 57,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "металлические",
//            "ROW_NUM": 58,
//            "OKEI_ID": 51,
//            "OKEI": "пог. м",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "из полимерных материалов",
//            "ROW_NUM": 59,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "из полимерных материалов",
//            "ROW_NUM": 60,
//            "OKEI_ID": 51,
//            "OKEI": "пог. м",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "деревянные",
//            "ROW_NUM": 61,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "деревянные",
//            "ROW_NUM": 62,
//            "OKEI_ID": 51,
//            "OKEI": "пог. м",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "Из строк 53 и 54 - трубы большого диаметра \n(2 и более метра) - всего",
//            "ROW_NUM": 63,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "Из строк 53 и 54 - трубы большого диаметра \n(2 и более метра) - всего",
//            "ROW_NUM": 64,
//            "OKEI_ID": 51,
//            "OKEI": "пог. м",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "2.5. Шумозащитные сооружения",
//            "ROW_NUM": 65,
//            "OKEI_ID": 796,
//            "OKEI": "пог. м",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "2.5. Шумозащитные сооружения",
//            "ROW_NUM": 66,
//            "OKEI_ID": 58,
//            "OKEI": "тыс. м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "2.6. Подпорные стены",
//            "ROW_NUM": 67,
//            "OKEI_ID": 18,
//            "OKEI": "пог. м",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "2.6. Подпорные стены",
//            "ROW_NUM": 68,
//            "OKEI_ID": 58,
//            "OKEI": "тыс. м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "2.7. Протяженность снегозаносимых \nучастков дорог",
//            "ROW_NUM": 69,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "2.8. Снегозащитные сооружения",
//            "ROW_NUM": 70,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "снегозащитные насаждения",
//            "ROW_NUM": 71,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "снегозащитные насаждения",
//            "ROW_NUM": 72,
//            "OKEI_ID": 59,
//            "OKEI": "га",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "снегозащитные средства \n(щиты, постоянные заборы)",
//            "ROW_NUM": 73,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "снегозащитные и противоселевые \nсооружения и конструкции \n(галереи, тоннели, лотки и пр.)",
//            "ROW_NUM": 74,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "3. ЗДАНИЯ \n  административные",
//            "ROW_NUM": 75,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "3. ЗДАНИЯ \n  административные",
//            "ROW_NUM": 76,
//            "OKEI_ID": 55,
//            "OKEI": "м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "жилищно-бытовые (жилые, гостиничные, \nстоловые и другие бытового назначения)",
//            "ROW_NUM": 77,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "жилищно-бытовые (жилые, гостиничные, \nстоловые и другие бытового назначения)",
//            "ROW_NUM": 78,
//            "OKEI_ID": 55,
//            "OKEI": "м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "производственные \n(мастерские, гаражи, склады и т.п.)",
//            "ROW_NUM": 79,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "производственные \n(мастерские, гаражи, склады и т.п.)",
//            "ROW_NUM": 80,
//            "OKEI_ID": 55,
//            "OKEI": "м2",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "базы и пункты для приготовления \nи хранения противогололедных средств \nи материалов",
//            "ROW_NUM": 81,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "4. ЗЕМЕЛЬНАЯ ПЛОЩАДЬ \nзанятая дорогами, дорожными  \nсооружениями и зданиями",
//            "ROW_NUM": 82,
//            "OKEI_ID": 59,
//            "OKEI": "га",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "5. ОБЪЕКТЫ АВТОГАЗОЗАПРАВОЧНОЙ \nИНФРАСТРУКТУРЫ \nАвтозаправочные станции (АЗС) - всего",
//            "ROW_NUM": 83,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "многотопливные заправочные станции \n(МТЗС)",
//            "ROW_NUM": 84,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "обеспечивающие заправку: сжиженным \nуглеводородным (нефтяным) газом \n(пропаном)",
//            "ROW_NUM": 85,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "компримированным природным газом \n(метаном)",
//            "ROW_NUM": 86,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "сжиженным природным газом \n(метаном)",
//            "ROW_NUM": 87,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "сжиженным, а также компривированным \nприродным газом (метаном)",
//            "ROW_NUM": 88,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "автомобильные газонаполнительные \nкомпрессорные станции (АГНКС)",
//            "ROW_NUM": 89,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "в том числе обеспечивающие заправку \nсжиженным природным газом (метаном)",
//            "ROW_NUM": 90,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "автомобильные газозаправочные станции \n(АГЗС)",
//            "ROW_NUM": 91,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "криогенные газозаправочные станции \n(КриоГЗС)",
//            "ROW_NUM": 92,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "в том числе также обеспечивающие \nкомпримированным природным газом \n(метаном)",
//            "ROW_NUM": 93,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "электрозаправочные станции (ЭЗС)",
//            "ROW_NUM": 94,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "BEG": 0,
//                "CONSTRUCTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "REPAIRED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "LISTED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "INVENTORIED": {
//                    "PLUS": 0,
//                    "MINUS": 0
//                },
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "- протяженность автомобильных дорог общего пользования, относящихся к \nфедеральной собственности (федерального значения) или к собственности \nсубъектов Российской Федерации (регионального или межмуниципального значения), в границах населенных пунктов",
//            "ROW_NUM": 95,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": " - протяженность автомобильных дорог общего пользования регионального или \nмежмуниципального значения, не отвечающих нормативным требованиям",
//            "ROW_NUM": 96,
//            "OKEI_ID": 8,
//            "OKEI": "км",
//            "VALUE": {
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "и их доля в общей протяженности автомобильных дорог \nобщего пользования регионального или межмуниципального значения",
//            "ROW_NUM": 97,
//            "OKEI_ID": 744,
//            "OKEI": "%",
//            "VALUE": {
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "- количество мостов, путепроводов, эстакад, \nнаходящихся в аварийном состоянии",
//            "ROW_NUM": 98,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "- количество мостов, путепроводов, эстакад, \nнаходящихся в предаварийном состоянии",
//            "ROW_NUM": 99,
//            "OKEI_ID": 796,
//            "OKEI": "шт",
//            "VALUE": {
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "- длина мостов, путепроводов, эстакад, находящихся в аварийном состоянии",
//            "ROW_NUM": 100,
//            "OKEI_ID": 18,
//            "OKEI": "пог.м",
//            "VALUE": {
//                "END": 0
//            }
//        },
//        {
//            "PARAM_NAME": "- длина мостов, путепроводов, эстакад, находящихся в предаварийном состоянии",
//            "ROW_NUM": 101,
//            "OKEI_ID": 18,
//            "OKEI": "пог.м",
//            "VALUE": {
//                "END": 0
//            }
//        }
//    ],
//    "FOOTER": {
//        "CAPTION": "",
//        "DATA": [
//            {
//                "FOOTER_ID": 1,
//                "COLUMN_NUM": "1",
//                "INDENT": 2,
//                "VALUE": "Ведущий консультант"
//            },
//            {
//                "FOOTER_ID": 1,
//                "COLUMN_NUM": "8",
//                "INDENT": 2,
//                "VALUE": "%EXECUTOR%"
//            }
//        ]
//    },
//    "STYLES": [
//        {
//            "GRID": [
//                1,
//                1,
//                1,
//                1
//            ],
//            "CELL_STYLE": {
//                "FONT_STYLE": {
//                    "BOLD": true
//                }
//            }
//        },
//        {
//            "GRID": [
//                2,
//                1,
//                2,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 2
//            }
//        },
//        {
//            "GRID": [
//                3,
//                1,
//                3,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 4
//            }
//        },
//        {
//            "GRID": [
//                4,
//                1,
//                6,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 8
//            }
//        },
//        {
//            "GRID": [
//                7,
//                1,
//                7,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 4
//            }
//        },
//        {
//            "GRID": [
//                8,
//                1,
//                9,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 6
//            }
//        },
//        {
//            "GRID": [
//                10,
//                1,
//                10,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 2
//            }
//        },
//        {
//            "GRID": [
//                11,
//                1,
//                11,
//                1
//            ],
//            "CELL_STYLE": {
//                "FONT_STYLE": {
//                    "BOLD": true
//                }
//            }
//        },
//        {
//            "GRID": [
//                12,
//                1,
//                12,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 2
//            }
//        },
//        {
//            "GRID": [
//                13,
//                1,
//                16,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 4
//            }
//        },
//        {
//            "GRID": [
//                17,
//                1,
//                17,
//                1
//            ],
//            "CELL_STYLE": {
//                "FONT_STYLE": {
//                    "BOLD": true
//                }
//            }
//        },
//        {
//            "GRID": [
//                18,
//                1,
//                18,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 2
//            }
//        },
//        {
//            "GRID": [
//                19,
//                1,
//                19,
//                1
//            ],
//            "CELL_STYLE": {
//                "FONT_STYLE": {
//                    "BOLD": true
//                }
//            }
//        },
//        {
//            "GRID": [
//                20,
//                1,
//                20,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 2
//            }
//        },
//        {
//            "GRID": [
//                21,
//                1,
//                22,
//                1
//            ],
//            "CELL_STYLE": {
//                "FONT_STYLE":{"BOLD": true}
//            }
//        },
//        {
//            "GRID": [
//                23,
//                1,
//                24,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 2
//            }
//        },
//        {
//            "GRID": [
//                25,
//                1,
//                25,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 2
//            }
//        },
//        {
//            "GRID": [
//                26,
//                1,
//                26,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 4
//            }
//        },
//        {
//            "GRID": [
//                27,
//                1,
//                27,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 4,
//                "FONT_STYLE":{"BOLD": true}
//            }
//        },
//        {
//            "GRID": [
//                28,
//                1,
//                28,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 4
//            }
//        },
//        {
//            "GRID": [
//                29,
//                1,
//                29,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 6
//            }
//        },
//        {
//            "GRID": [
//                30,
//                1,
//                30,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 8
//            }
//        },
//        {
//            "GRID": [
//                31,
//                1,
//                32,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 4,
//                "FONT_STYLE": {
//                    "BOLD": true
//                }
//            }
//        },
//        {
//            "GRID": [
//                33,
//                1,
//                33,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 6
//            }
//        },
//        {
//            "GRID": [
//                34,
//                1,
//                34,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 8
//            }
//        },
//        {
//            "GRID": [
//                35,
//                1,
//                36,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 4,
//                "FONT_STYLE": {
//                    "BOLD": true
//                }
//            }
//        },
//        {
//            "GRID": [
//                37,
//                1,
//                37,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 6
//            }
//        },
//        {
//            "GRID": [
//                38,
//                1,
//                38,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 8
//            }
//        },
//        {
//            "GRID": [
//                39,
//                1,
//                39,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 4,
//                "FONT_STYLE": {
//                    "BOLD": true
//                }
//            }
//        },
//        {
//            "GRID": [
//                41,
//                1,
//                42,
//                1
//            ],
//            "CELL_STYLE": {
//                "FONT_STYLE": {
//                    "BOLD": true
//                }
//            }
//        },
//        {
//            "GRID": [
//                43,
//                1,
//                43,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 2
//            }
//        },
//        {
//            "GRID": [
//                44,
//                1,
//                44,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 4
//            }
//        },
//        {
//            "GRID": [
//                45,
//                1,
//                46,
//                1
//            ],
//            "CELL_STYLE": {
//                "FONT_STYLE": {
//                    "BOLD": true
//                }
//            }
//        },
//        {
//            "GRID": [
//                47,
//                1,
//                48,
//                1
//            ],
//            "CELL_STYLE": {
//                "FONT_STYLE": {
//                    "BOLD": true
//                },
//            "INDENT":2
//            }
//        },
//        {
//            "GRID": [
//                49,
//                1,
//                49,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 4
//            }
//        },
//        {
//            "GRID": [
//                50,
//                1,
//                51,
//                1
//            ],
//            "CELL_STYLE": {
//                "FONT_STYLE": {
//                    "BOLD": true
//                },
//                "INDENT":2
//            }
//        },
//        {
//            "GRID": [
//                52,
//                1,
//                52,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 4
//            }
//        },
//        {
//            "GRID": [
//                53,
//                1,
//                54,
//                1
//            ],
//            "CELL_STYLE": {
//                "FONT_STYLE": {
//                    "BOLD": true
//                }
//            }
//        },
//        {
//            "GRID": [
//                55,
//                1,
//                62,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 2
//            }
//        },
//        {
//            "GRID": [
//                63,
//                1,
//                64,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 0
//            }
//        },
//        {
//            "GRID": [
//                65,
//                1,
//                70,
//                1
//            ],
//            "CELL_STYLE": {
//                "FONT_STYLE": {
//                    "BOLD": true
//                }
//            }
//        },
//        {
//            "GRID": [
//                71,
//                1,
//                74,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 2
//            }
//        },
//        {
//            "GRID": [
//                75,
//                1,
//                76,
//                1
//            ],
//            "CELL_STYLE": {
//                "FONT_STYLE":{"BOLD":true}
//            }
//        },
//        {
//            "GRID": [
//                77,
//                1,
//                81,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT":2
//            }
//        },
//        {
//            "GRID": [
//                82,
//                1,
//                83,
//                1
//            ],
//            "CELL_STYLE": {
//                "FONT_STYLE":{"BOLD":true}
//            }
//        },
//        {
//            "GRID": [
//                84,
//                1,
//                84,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 2
//            }
//        },
//        {
//            "GRID": [
//                85,
//                1,
//                88,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 4
//            }
//        },
//        {
//            "GRID": [
//                89,
//                1,
//                89,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 2
//            }
//        },
//        {
//            "GRID": [
//                90,
//                1,
//                90,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 2
//            }
//        },
//        {
//            "GRID": [
//                91,
//                1,
//                92,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 2
//            }
//        },
//        {
//            "GRID": [
//                93,
//                1,
//                93,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 4
//            }
//        },
//        {
//            "GRID": [
//                94,
//                1,
//                94,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 2
//            }
//        },
//        {
//            "GRID": [
//                95,
//                1,
//                97,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 2
//            }
//        },
//        {
//            "GRID": [
//                98,
//                1,
//                99,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 2
//            }
//        },
//        {
//            "GRID": [
//                100,
//                1,
//                101,
//                1
//            ],
//            "CELL_STYLE": {
//                "INDENT": 2
//            }
//        },
//        {
//            "GRID": [
//                83,
//                5,
//                88,
//                12
//            ],
//            "PROTECTED": true
//        }
//    ]
//}
//`)
