package httpdtoreq

type HttpFindIncidentsDto struct {
	Description string `form:"description" binding:"omitempty,max=50"`
	Limit       int    `form:"limit" binding:"min=1,max=100"`
	Offset      int    `form:"offset" binding:"omitempty,min=0"`
}

func (dto *HttpFindIncidentsDto) SetDefaults() {
	if dto.Limit == 0 {
		dto.Limit = 10
	}
}
