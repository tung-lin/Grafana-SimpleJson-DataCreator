package main

import (
	"bytes"
	"fmt"
	"grafana-simplejson-datacreator/common/db"
	"grafana-simplejson-datacreator/common/dto"
	"strings"
)

type APIProviderLookup struct {
}

var DataCreator APIProviderLookup

func main() {

}

func (e APIProviderLookup) GetIdentifyName() string {
	return "apiproviderlookup"
}

func (e APIProviderLookup) CreateData(keys []string) []dto.SearchResult {

	results := []dto.SearchResult{}

	if len(keys) == 0 {

		rows, error := db.GetDBData("MOTCAPI_v2", "select ID from Provider order by NameZh_tw")
		defer rows.Close()

		if error == nil {
			for rows.Next() {
				var id string
				rows.Scan(&id)
				results = append(results, dto.SearchResult{Text: id, Value: id})
			}
		}

	} else {

		for index, key := range keys {
			keyPos := &keys[index]
			*keyPos = fmt.Sprintf("'%s'", key)
		}

		var buffer bytes.Buffer
		buffer.WriteString("(")
		buffer.WriteString(strings.Join(keys, ","))
		buffer.WriteString(")")

		rows, error := db.GetDBData("MOTCAPI_v2", "select ID, NameZh_tw from Provider where ID in "+buffer.String()+" order by NameZh_tw")
		defer rows.Close()

		if error == nil {
			for rows.Next() {
				var id string
				var name string
				rows.Scan(&id, &name)
				results = append(results, dto.SearchResult{Text: name, Value: id})
			}
		}
	}

	return results
}
