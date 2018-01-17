package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"grafana-simplejson-datacreator/common/db"
	"grafana-simplejson-datacreator/common/dto"
	"log"
	"strings"
)

type UserLookup struct {
}

var DataCreator UserLookup

func (e UserLookup) GetIdentifyName() string {
	return "userlookup"
}

func (e UserLookup) CreateData(keys []string) []dto.SearchResult {

	results := []dto.SearchResult{}

	if len(keys) == 0 {
		rows, error := db.GetDBData("MOTCWeb_v2", "select PK_User from user where FK_UserStatus = 'Status_Able' order by AppName desc")
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

		rows, error := db.GetDBData("MOTCWeb_v2", "select PK_User, OrganizationName, APPName from user where FK_UserStatus = 'Status_Able' and PK_User in "+buffer.String()+" order by AppName desc")
		defer rows.Close()

		if error == nil {
			for rows.Next() {
				var id string
				var organizationname sql.NullString
				var appname sql.NullString

				if err := rows.Scan(&id, &organizationname, &appname); err != nil {
					log.Printf("an error occurred while scaning rows\r\n%s", err)
				} else {
					results = append(results, dto.SearchResult{Value: id, Text: appname.String + "(" + organizationname.String + ")"})
				}
			}
		}
	}

	return results
}
