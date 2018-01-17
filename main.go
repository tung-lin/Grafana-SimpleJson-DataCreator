package main

import (
	"encoding/json"
	"fmt"
	"grafana-simplejson-datacreator/common/config"
	"grafana-simplejson-datacreator/common/dto"
	"grafana-simplejson-datacreator/common/plugininterface"
	"log"
	"net/http"
	"os"
	"plugin"
	"strings"

	"github.com/gorilla/mux"
)

func init() {

}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", action_default).Methods("GET")
	router.HandleFunc("/search", action_search).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+config.Config.Port, router))
}

func action_default(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Just for test"))
	//json.NewEncoder(w).Encode("Just for test")
}

func action_search(w http.ResponseWriter, r *http.Request) {
	var searchData dto.SearchData
	var searchDetailData dto.SearchDetailData

	json.NewDecoder(r.Body).Decode(&searchData)
	json.Unmarshal([]byte(searchData.Target), &searchDetailData)

	//log.Println(searchDetailData.Data)
	//log.Println(searchDetailData.ID)

	var results []dto.SearchResult
	creator := getPlugin(searchDetailData.Data)

	if creator != nil {
		var ids []string

		if searchDetailData.ID != "" {
			replacer := strings.NewReplacer("(", "", ")", "")
			searchDetailData.ID = replacer.Replace(searchDetailData.ID)
			ids = strings.Split(searchDetailData.ID, "|")
		}

		results = creator.CreateData(ids)
		json.NewEncoder(w).Encode(results)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("data '%s' not found!", searchDetailData.Data)))
	}
}

func getPlugin(pluginName string) plugininterface.IDataCreator {

	mainPath, _ := os.Getwd()
	pluginPath := fmt.Sprintf("%s/datacreator/%s/%s.so", mainPath, pluginName, pluginName)
	plugin, err := plugin.Open(pluginPath)

	if err != nil {
		log.Printf("load symbol file '%s' failed: %s", pluginName, err)
		return nil
	}

	symbol, err := plugin.Lookup("DataCreator")

	if err != nil {
		log.Printf("look up a symbol 'DataCreator' failed: %s", err)
		return nil
	}

	var creator plugininterface.IDataCreator
	creator, ok := symbol.(plugininterface.IDataCreator)

	if !ok {
		log.Printf("assert loaded symbol to type 'IDataCreator' failed: %s", err)
		return nil
	}

	return creator
}
