package csvmapping

import "StaticData/domain"

type Repository interface {
	// blbubb
	AllByKind(kind string) []*domain.CsvMapping

	//tjrejejej
	AllByKindAndParticipant(kind string, participant int) []*domain.CsvMapping
}
