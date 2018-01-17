package plugininterface

import "grafana-simplejson-datacreator/common/dto"

type IDataCreator interface {
	GetIdentifyName() string
	CreateData(keys []string) []dto.SearchResult
}
