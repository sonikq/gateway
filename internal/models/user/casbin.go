package user

type Subject struct {
	UserID          int
	ObjectValue     int64
	ObjectType      string
	PeriodStart     string
	PeriodFinish    string
	States          []string
	Role            string
	OrganizationsID []int64
}

type Object struct {
	OKUD           int64
	PeriodStart    string
	PeriodFinish   string
	DocState       string
	OrgResponderID int64
}
