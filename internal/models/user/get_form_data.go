package user

import "gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"

type GetFormDataRequest struct {
	UserID int
	FormID int
}

type GetFormDataResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response *FormData
}

type Form struct {
	FormID       int
	ChapterIDMap map[int]string
}

type GetFormIDByChapterID struct {
	FormID int `json:"form_id"`
}

type FormData struct {
	DOCS []struct {
		ERROR   bool   `json:"ERROR,omitempty"`
		ID      int    `json:"ID,omitempty"`
		TITLE   string `json:"TITLE,omitempty"`
		OKUD    string `json:"OKUD,omitempty"`
		CODE    string `json:"CODE,omitempty"`
		LABEL   string `json:"LABEL,omitempty"`
		CUSTOM  bool   `json:"CUSTOM,omitempty"`
		CHAPTER struct {
			TEXT  string `json:"TEXT,omitempty"`
			STYLE STYLE  `json:"STYLE,omitempty"`
		} `json:"CHAPTER,omitempty"`
		PERIOD         string `json:"PERIOD,omitempty"`
		ORGRESPONDERID int    `json:"ORG_RESPONDER_ID,omitempty"`
		ORGRESPONDER   string `json:"ORG_RESPONDER,omitempty"`
		ORGREVIEWERID  int    `json:"ORG_REVIEWER_ID,omitempty"`
		EXECUTORID     int    `json:"EXECUTOR_ID,omitempty"`
		EXECUTORFIO    string `json:"EXECUTOR_FIO,omitempty"`
		REGION         string `json:"REGION,omitempty"`
		REGIONOKATO    string `json:"REGION_OKATO,omitempty"`
		REGIONID       int    `json:"REGION_ID,omitempty"`
		PREHEADER      struct {
			TEXT  string `json:"TEXT,omitempty"`
			STYLE STYLE  `json:"STYLE,omitempty"`
		} `json:"PRE_HEADER,omitempty"`
		HEADER    HEADER     `json:"HEADER,omitempty"`
		COMMON    COMMON     `json:"COMMON,omitempty"`
		COLUMNS   []COLUMN   `json:"COLUMNS"`
		ROWUNIONS []ROWUNION `json:"ROW_UNIONS,omitempty"`
		ROWBREAKS []ROWBREAK `json:"ROW_BREAKS,omitempty"`
		READONLY  [][]int    `json:"READONLY"`
		STYLES    []struct {
			GRID       []int `json:"GRID,omitempty"`
			CELL_STYLE STYLE `json:"CELL_STYLE,omitempty"`
			PROTECTED  bool  `json:"PROTECTED,omitempty"`
		} `json:"STYLES,omitempty"`
		DATA   []map[string]interface{} `json:"DATA"`
		FOOTER FOOTER                   `json:"FOOTER,omitempty"`
	} `json:"DOCS"`
	TITLEPAGE struct {
		OKPO      string `json:"OKPO"`
		OKUD      string `json:"OKUD"`
		TITLE     string `json:"TITLE"`
		FORMNAME  string `json:"FORM_NAME"`
		ORDERNUM  string `json:"ORDER_NUM"`
		RESPONSES []struct {
			DEADLINE     string   `json:"DEADLINE"`
			ORGRESPONDER string   `json:"ORG_RESPONDER"`
			ORGREVIEWERS []string `json:"ORG_REVIEWERS"`
		} `json:"RESPONSES"`
		ORDERDATE    string `json:"ORDER_DATE"`
		REPORTTYPE   string `json:"REPORT_TYPE"`
		ORGANIZATION string `json:"ORGANIZATION"`
		STATEBYDATE  string `json:"STATE_BY_DATE"`
	} `json:"TITLE_PAGE"`
}
