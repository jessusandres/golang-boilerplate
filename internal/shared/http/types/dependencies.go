package types

import incidentsinterfaces "lookerdevelopers/boilerplate/internal/modules/incident/interfaces"

type RouterDependencies struct {
	IncidentController incidentsinterfaces.IIncidentsController
}
