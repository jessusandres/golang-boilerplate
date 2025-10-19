package queries

type FindIncidentsQuery struct {
	Description string
	Limit       int
	Offset      int
}

func (c FindIncidentsQuery) QueryName() string {
	return "find_incidents"
}
